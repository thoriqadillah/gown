package chunk

import (
	"bufio"
	_http "changeme/gown/http"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Chunk struct {
	response _http.Response
	wg       *sync.WaitGroup
	index    int
	start    int64
	end      int64
	size     int64
	data     []byte
	ctx      context.Context
}

func New(ctx context.Context, res _http.Response, index int, wg *sync.WaitGroup) *Chunk {
	totalpart := int64(res.Totalpart)
	partsize := res.Size / totalpart

	start := int64(index * int(partsize))
	end := start + int64(int(partsize)-1)

	if index == int(totalpart)-1 {
		end = res.Size
	}

	return &Chunk{
		response: res,
		wg:       wg,
		index:    index,
		start:    start,
		end:      end,
		size:     partsize,
		data:     make([]byte, partsize),
		ctx:      ctx,
	}
}

func (c *Chunk) download() error {
	defer c.wg.Done()

	//TODO: implement download chunk
	http_ := &http.Client{}
	part := fmt.Sprintf("bytes=%d-%d", c.start, c.end)

	if c.size == -1 {
		log.Printf("Downloading chunk %d with size unknown", c.index+1)
	} else {
		log.Printf("Downloading chunk %d from %d to %d (~%d MB)", c.index+1, c.start, c.end, (len(c.data))/(1024*1024))
	}

	req, err := http.NewRequest("GET", c.response.Url, nil)
	if err != nil {
		return err
	}

	start := time.Now()

	req.Header.Add("Range", part)
	res, err := http_.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var _100KB int64 = 1024 * 100
	r := bufio.NewReader(res.Body)
	var transfered int64

	for {
		transfered += _100KB
		if transfered > c.size {
			_100KB = transfered - c.size
		}
		buffer := make([]byte, _100KB)

		n, err := io.ReadFull(r, buffer)
		if err == io.EOF {
			break
		}

		//TODO: combine the downloaded bytes into c.data
		c.data = append(c.data, buffer...)
		runtime.EventsEmit(c.ctx, "transfered", c.index, n)
	}

	// c.data, err = io.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }

	// runtime.EventsEmit(c.ctx, "transfered", len(c.data))

	elapsed := time.Since(start)
	log.Printf("Chunk %d downloaded in %v s\n", c.index+1, elapsed.Seconds())

	return nil
}

func (c *Chunk) Execute() error {
	var err error

	for retry := 0; retry < c.response.Settings.Maxtries; retry++ {
		if err = c.download(); err == nil {
			return nil
		}

		log.Printf("Error while downloading chunk %d: %v. Retrying....\n", c.index+1, err)
	}

	return err
}

// TODO: implement handle error
func (c *Chunk) HandleError(err error) {
	log.Printf("Error while downloading chunk %d: %v\n", c.index+1, err)
}

func (c Chunk) Data() []byte {
	return c.data
}
