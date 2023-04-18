package chunk

import (
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
		data:     make([]byte, 0, partsize),
		ctx:      ctx,
	}
}

func (c *Chunk) download() error {
	defer c.wg.Done()

	http_ := &http.Client{}
	part := fmt.Sprintf("bytes=%d-%d", c.start, c.end)

	if c.size == -1 {
		log.Printf("Downloading chunk %d with size unknown", c.index+1)
	} else {
		log.Printf("Downloading chunk %d from %d to %d (~%d MB)", c.index+1, c.start, c.end, (c.size)/(1024*1024))
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
	buffer := make([]byte, _100KB)

	for {
		n, err := io.ReadFull(res.Body, buffer)
		if err == io.EOF {
			break
		}

		c.data = append(c.data, buffer[:n]...)
		runtime.EventsEmit(c.ctx, "transfered", c.index, n) //TODO: implement proper transfer emition to differentiate which div to animate the progress bar
	}

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
