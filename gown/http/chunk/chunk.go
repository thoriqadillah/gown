package chunk

import (
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Chunk struct {
	toDownload *download.Download
	wg         *sync.WaitGroup
	index      int
	start      int64
	end        int64
	size       int64
	ctx        context.Context
	*setting.Settings
}

func New(ctx context.Context, toDownload *download.Download, index int, setting *setting.Settings, wg *sync.WaitGroup) *Chunk {
	totalpart := int64(toDownload.Metadata.Totalpart)
	partsize := toDownload.Size / totalpart

	start := int64(index * int(partsize))
	end := start + int64(int(partsize)-1)

	if index == int(totalpart)-1 {
		end = toDownload.Size
	}

	return &Chunk{
		toDownload: toDownload,
		wg:         wg,
		index:      index,
		start:      start,
		end:        end,
		size:       partsize,
		ctx:        ctx,
		Settings:   setting,
	}
}

func (c *Chunk) download() error {
	http_ := &http.Client{}
	part := fmt.Sprintf("bytes=%d-%d", c.start, c.end)

	if c.size == -1 {
		log.Printf("Downloading chunk %d with size unknown", c.index+1)
	} else {
		log.Printf("Downloading chunk %d from %d to %d (~%d MB)", c.index+1, c.start, c.end, (c.end-c.start)/(1024*1024))
	}

	req, err := http.NewRequest("GET", c.toDownload.Metadata.Url, nil)
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

	// create temp file
	tmpFilename := filepath.Join(c.SaveLocation, fmt.Sprintf("%s-%d", c.toDownload.ID, c.index))
	file, err := os.OpenFile(tmpFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error creating or appending file: %v\n", err)
		return err
	}
	defer file.Close()

	progressbar := &progressbar{
		ctx:        c.ctx,
		toDownload: c.toDownload,
		index:      c.index,
		Reader:     res.Body,
		partsize:   c.size,
	}

	if _, err := io.Copy(file, progressbar); err != nil {
		return err
	}

	elapsed := time.Since(start)
	log.Printf("Chunk %d downloaded in %v s\n", c.index+1, elapsed.Seconds())

	if err == nil {
		c.wg.Done()
	}
	return nil
}

func (c *Chunk) ResumeFrom(position int64) *Chunk {
	c.start += position
	return c
}

func (c *Chunk) Execute() error {
	var err error

	for retry := 0; retry < c.Settings.Maxtries; retry++ {
		if err = c.download(); err == nil {
			return nil
		}

		if err == errCanceled {
			return err
		}

		log.Printf("Error while downloading chunk %d: %v. Retrying....\n", c.index+1, err)
	}

	return err
}

// TODO: implement handle error
func (c *Chunk) HandleError(err error) {
	defer c.wg.Done()

	if err == errCanceled {
		log.Println("download canceled")
		//TODO: save the downloaded data and mark the range if resumable. otherwise, delete the temp file
	} else {
		// TODO: retry from the position where error happened
		log.Printf("Error while downloading chunk %d: %v\n", c.index+1, err)
	}

}
