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
	"strconv"
	"sync"
	"time"
)

type Chunk struct {
	toDownload download.Download
	wg         *sync.WaitGroup
	index      int
	start      int64
	end        int64
	size       int64
	ctx        context.Context
	tmpFile    *os.File
	*setting.Settings
}

func New(ctx context.Context, toDownload download.Download, index int, setting *setting.Settings, wg *sync.WaitGroup) *Chunk {
	totalpart := int64(toDownload.Metadata.Totalpart)
	partsize := toDownload.Size / totalpart

	start := int64(index * int(partsize))
	end := start + int64(int(partsize)-1)

	if index == int(totalpart)-1 {
		end = toDownload.Size
	}

	tmp := filepath.Join(setting.SaveLocation, toDownload.ID+"-"+strconv.Itoa(index))
	f, err := os.Create(tmp)
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
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
		tmpFile:    f,
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

	progressbar := &progressbar{
		ctx:    c.ctx,
		id:     c.toDownload.ID,
		index:  c.index,
		Reader: res.Body,
		size:   c.size,
		tmp:    0,
	}

	if _, err := io.Copy(c.tmpFile, progressbar); err != nil {
		// TODO: retry from the position where error happened
		return err
	}

	elapsed := time.Since(start)
	log.Printf("Chunk %d downloaded in %v s\n", c.index+1, elapsed.Seconds())

	defer c.tmpFile.Close()
	return nil
}

func (c *Chunk) Execute() error {
	var err error

	for retry := 0; retry < c.Settings.Maxtries; retry++ {
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
