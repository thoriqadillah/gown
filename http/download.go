package http

import (
	"sync"

	"github.com/thoriqadillah/gown/worker"
)

type download struct {
	wg *sync.WaitGroup
	chunk
}

type chunk struct {
	index int
	data  []byte
}

func Download(file *response, index int, wg *sync.WaitGroup) worker.Job {
	// get the part size that we want to download
	totalpart := int64(file.totalpart)
	partsize := file.size / totalpart
	if index == file.totalpart {
		partsize = file.size - ((totalpart - 1) * file.size / totalpart) // size of the last index of the chunk is the remaining bytes
	}

	chuck := chunk{
		index: index,
		data:  make([]byte, partsize),
	}

	return &download{
		wg:    wg,
		chunk: chuck,
	}
}

func (d *download) Execute() error {
	defer d.wg.Done()

	// TODO: implement ranged download

	return nil
}

func (d *download) HandleError(err error) {
	// TODO: handle error
}
