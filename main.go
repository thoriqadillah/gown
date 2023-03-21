package main

import (
	"log"
	"sync"

	"github.com/thoriqadillah/gown/http"
	pool "github.com/thoriqadillah/gown/worker"
)

func main() {
	const TOTAL_WORKER = 8
	const TOTAL_POOL = 1

	worker, err := pool.New(TOTAL_WORKER, TOTAL_POOL)
	if err != nil {
		log.Fatal("Error creating worker")
	}
	worker.Start()
	defer worker.Stop()

	var wg sync.WaitGroup

	response, err := http.Fetch("https://get.jenkins.io/war-stable/2.332.3/jenkins.war")
	if err != nil {
		log.Fatal(err)
	}

	downloadjob := make([]pool.Job, response.Parts())
	for part := range downloadjob {
		downloadjob[part] = http.Download(response, part, &wg)
	}

	for _, job := range downloadjob {
		wg.Add(1)
		worker.Add(job)
	}

	wg.Wait()
}
