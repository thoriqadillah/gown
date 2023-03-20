package http

import (
	"sync"

	"github.com/thoriqadillah/gown/worker"
)

type Partition struct {
	wg *sync.WaitGroup
	chunk
}

type chunk struct {
	index int
	data  []byte
}

func Download(file *response, index int, wg *sync.WaitGroup) worker.Job {
	// get the part size that we want to download
	partsize := file.size / int64(file.totalpart)
	if index == file.totalpart {
		partsize = file.size - (7 * file.size / 8)
	}

	chuck := chunk{
		index: index,
		data:  make([]byte, partsize),
	}

	return &Partition{
		wg:    wg,
		chunk: chuck,
	}
}

func (d *Partition) Execute() error {
	defer d.wg.Done()

	return nil
}

func (d *Partition) HandleError(err error) {

}
