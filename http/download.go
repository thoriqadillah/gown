package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/thoriqadillah/gown/config"
	"github.com/thoriqadillah/gown/worker"
)

type Chunk struct {
	*config.Config
	*response
	index int
	start int64
	end   int64
	Data  []byte
}

func Download(res *response, index int, conf ...*config.Config) worker.Job {
	cfg := config.Default()
	if conf != nil {
		cfg = conf[0]
	}

	totalpart := res.size / cfg.Partsize
	start := int64(index * int(cfg.Partsize))
	end := start + cfg.Partsize - 1

	if index == int(totalpart)-1 {
		end = res.size
	}

	return &Chunk{
		Config:   cfg,
		response: res,
		index:    index,
		start:    start,
		end:      end,
		Data:     []byte{},
	}
}

func (c *Chunk) download() error {
	start := time.Now()
	httpclient := &http.Client{}

	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		log.Printf("Error while making request: %v\n", err)
		return err
	}

	part := fmt.Sprintf("bytes=%d-%d", c.start, c.end)
	req.Header.Add("Range", part)

	log.Printf("Downloading chunk %d\n", c.index+1)
	res, err := httpclient.Do(req)
	if err != nil {
		log.Printf("Error while downloading chunk %d: %v\n", c.index+1, err)
		return err
	}

	c.Data, err = io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error while downloading chunk %d: %v\n", c.index+1, err)
		return err
	}

	elapsed := time.Since(start)
	log.Printf("Chunk %d downloaded in %v s", c.index+1, elapsed.Seconds())

	return nil
}

func (c *Chunk) Execute() error {
	var err error
	for retry := 0; retry < c.Maxtries; retry++ {
		if err = c.download(); err == nil {
			return nil
		}

		log.Printf("Error while downloading chunk %d: %v . Retrying\n", c.index+1, err)
	}

	return err
}

func (c *Chunk) HandleError(err error) {
	log.Printf("Error while downloading the chunk %d: %v\n", c.index+1, err)
}

func (c *Chunk) Struct() interface{} {
	return c
}
