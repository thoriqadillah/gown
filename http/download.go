package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thoriqadillah/gown/worker"
)

type Partition struct {
	wg       *sync.WaitGroup
	url      string
	callback func(data []byte, index int)
	chunk
}

type chunk struct {
	index int
	start int64
	end   int64
	size  int64
	Data  []byte
}

func Download(res *response, index int, wg *sync.WaitGroup, callback func(data []byte, index int)) worker.Job {
	// get the range part that we want to download
	totalpart := int64(res.totalpart)
	partsize := res.size / totalpart

	start := int64(index * int(partsize))
	end := start + int64(int(partsize)-1)

	if index == int(totalpart)-1 {
		end = res.size
	}

	chuck := chunk{
		index: index,
		start: start,
		end:   end,
		size:  partsize,
	}

	return &Partition{
		wg:       wg,
		url:      res.url,
		callback: callback,
		chunk:    chuck,
	}
}

func (d *Partition) Execute() error {
	start := time.Now()
	defer d.wg.Done()

	httpclient := &http.Client{}

	part := fmt.Sprintf("bytes=%d-%d", d.start, d.end)

	if d.size == -1 {
		log.Printf("Downloading part %d with size unknown", d.index+1)
	} else {
		log.Printf("Downloading part %d from %d to %d", d.index+1, d.start, d.end)
	}

	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", part)
	res, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// callback to combine the file
	d.callback(data, d.index)
	elapsed := time.Since(start)
	fmt.Printf("Worker with id %d done in %v s\n", d.index, elapsed.Seconds())
	return nil
}

func (d *Partition) HandleError(err error) {
	// TODO: handle error
	log.Println(err)
}
