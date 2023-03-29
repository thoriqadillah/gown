package worker

import (
	"fmt"
	"log"
	"sync"
)

type Pool interface {
	Start()
	Stop()
	Add(Job)
}

type Job interface {
	Execute() error
	HandleError(error)
	Struct() interface{}
}

type worker struct {
	amount int
	jobs   chan Job
	start  sync.Once
	stop   sync.Once
	quit   chan struct{}
}

func New(amount int, poolsize int) (Pool, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("worker cannot less than 1")
	}

	if poolsize <= 0 {
		return nil, fmt.Errorf("worker pool canot be less than 1")
	}

	worker := worker{
		amount: amount,
		jobs:   make(chan Job, poolsize),
		start:  sync.Once{},
		stop:   sync.Once{},
		quit:   make(chan struct{}),
	}

	return &worker, nil
}

// Create and start n amount of workers to read from the jobs channel and process the job
func (w *worker) Start() {
	w.start.Do(func() {
		log.Println("Starting the worker")

		for i := 0; i < w.amount; i++ {
			go func(id int) {
				log.Printf("Starting worker with id %d\n", id)

				for {
					select {
					case <-w.quit:
						log.Printf("Stopping worker %d with quit channel\n", id)
						return
					case job, ok := <-w.jobs:
						if !ok {
							log.Printf("Stopping worker %d with closed channel\n", id)
							return
						}

						if err := job.Execute(); err != nil {
							job.HandleError(err)
						}
					}
				}
			}(i)
		}
	})
}

// Stop the worker from working
func (w *worker) Stop() {
	w.stop.Do(func() {
		log.Println("Stopping worker")
		close(w.quit)
		close(w.jobs)
	})
}

// Add work to the worker. If the channel buffer is full (or 0) and
// all workers are occupied, this will hang until work is consumed or Stop() is called.
func (w *worker) Add(job Job) {
	select {
	case w.jobs <- job:
	case <-w.quit:
	}
}
