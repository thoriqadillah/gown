package worker

import (
	"fmt"
	"log"
	"sync"
)

type Worker struct {
	amount int
	jobs   chan Job
	start  sync.Once
	stop   sync.Once
	quit   chan struct{}
}

func New(amount int, poolsize int) (Pool, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("invalid worker amount")
	}

	if poolsize <= 0 {
		return nil, fmt.Errorf("invalid pool size")
	}

	return &Worker{
		amount: amount,
		jobs:   make(chan Job, poolsize),
		start:  sync.Once{},
		stop:   sync.Once{},
		quit:   make(chan struct{}),
	}, nil
}

// Create and start n amount of workers to read from the jobs channel and process the job
func (w *Worker) Start() {
	w.start.Do(func() {
		log.Printf("Starting %d workers\n", w.amount)

		for i := 0; i < w.amount; i++ {
			go func(id int) {
				log.Printf("Starting worker %d\n", id)

				for {
					select {
					case <-w.quit:
						log.Printf("Worker %d stopped from quit channel\n", id)
						return

					case job, ok := <-w.jobs:
						if !ok {
							log.Printf("Worker %d stopped from jobs channel\n", id)
							return
						}

						if err := job.Execute(); err != nil {
							log.Printf("Worker %d failed to execute job: %s\n", id, err)
							job.HandleError(err)
						}
					}
				}
			}(i)
		}
	})
}

func (w *Worker) Stop() {
	w.stop.Do(func() {
		log.Printf("Stopping %d workers\n", w.amount)
		close(w.quit)
		close(w.jobs)
	})
}

// Add work to the worker. If the channel buffer is full (or 0) and
// all workers are occupied, this will hang until work is consumed or Stop() is called.
func (w *Worker) Add(job Job) {
	select {
	case w.jobs <- job:
	case <-w.quit:
	}
}
