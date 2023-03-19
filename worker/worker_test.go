package worker

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {
	if _, err := New(0, 0); err == nil {
		t.Fatalf("expected error when creating 0 worker, got %v", err)
	}

	if _, err := New(-1, 0); err == nil {
		t.Fatalf("expected error when creating -1 worker, got %v", err)
	}

	if _, err := New(1, -1); err == nil {
		t.Fatalf("expected error when creating -1 channel, got %v", err)
	}

	worker, err := New(5, 0)
	if worker == nil {
		t.Fatalf("worker returned nil with valid input")
	}

	if err != nil {
		t.Fatalf("expected no error when creating valid worker, got %v", err)
	}
}

func TestMultipleStartAndStop(t *testing.T) {
	worker, err := New(5, 0)
	if err != nil {
		t.Fatalf("expected no error when creating valid worker, got %v", err)
	}

	// Checking to make sure multiple calls to start or stop don't cause a panic
	worker.Start()
	worker.Start()

	worker.Stop()
	worker.Stop()
}

type testJob struct {
	executeFunc func() error

	shouldErr bool
	wg        *sync.WaitGroup

	mFailure       *sync.Mutex
	failureHandled bool
}

func newTestJob(executeFunc func() error, shouldErr bool, wg *sync.WaitGroup) *testJob {
	return &testJob{
		executeFunc: executeFunc,
		shouldErr:   shouldErr,
		wg:          wg,
		mFailure:    &sync.Mutex{},
	}
}

func (t *testJob) Execute() error {
	if t.wg != nil {
		defer t.wg.Done()
	}

	if t.executeFunc != nil {
		return t.executeFunc()
	}

	// if no function provided, just wait and error if told to do so
	time.Sleep(50 * time.Millisecond)
	if t.shouldErr {
		return fmt.Errorf("planned Execute() error")
	}
	return nil
}

func (t *testJob) HandleError(e error) {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()

	t.failureHandled = true
}

func (t *testJob) hitFailureCase() bool {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()

	return t.failureHandled
}

func TestWorkerPool_Work(t *testing.T) {
	var jobs []*testJob
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		jobs = append(jobs, newTestJob(nil, false, wg))
	}

	worker, err := New(5, len(jobs))
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	worker.Start()

	for _, job := range jobs {
		worker.Add(job)
	}

	// we'll get a timeout failure if the jobs weren't processed
	wg.Wait()

	for jobNum, job := range jobs {
		if job.hitFailureCase() {
			t.Fatalf("error function called on job %d when it shouldn't be", jobNum)
		}
	}
}

func TestWorkerPool_BlockedAddWorkReleaseAfterStop(t *testing.T) {
	worker, err := New(1, 1)
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}

	worker.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		// the first should start processing right away, the second two should hang
		wg.Add(1)
		go func() {
			worker.Add(newTestJob(func() error {
				time.Sleep(20 * time.Second)
				return nil
			}, false, nil))
			wg.Done()
		}()
	}

	done := make(chan struct{})
	worker.Stop()
	go func() {
		// wait on our Add() calls to complete, then signal on the done channel
		wg.Wait()
		done <- struct{}{}
	}()

	// wait until either we hit our timeout, or we're told the AddWork calls completed
	select {
	case <-time.After(1 * time.Second):
		t.Fatal("failed because still hanging on Add()")
	case <-done: // this is the success case
	}
}
