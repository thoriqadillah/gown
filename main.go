package main

import (
	"sync"

	"github.com/thoriqadillah/gown/http/api"
	"github.com/thoriqadillah/gown/http/service"
	"github.com/thoriqadillah/gown/worker"
)

func main() {
	worker, err := worker.New(8, 1)
	if err != nil {
		panic(err)
	}
	worker.Start()
	defer worker.Stop()

	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"
	years := []string{"2013", "2012", "2011", "2010", "2009", "2008", "2007", "2001"}
	api := api.New(url)

	var wg sync.WaitGroup
	for i := range years {
		wg.Add(1)
		job := service.NewGraduation(api, years[i], &wg)
		worker.Add(job)
	}

	wg.Wait()
}
