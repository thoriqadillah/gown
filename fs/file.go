package fs

import (
	"os"
	"path/filepath"
	"sync"
)

type File struct {
	Data  [][]byte
	mutex *sync.Mutex
}

func New(size int64) *File {
	return &File{
		Data:  make([][]byte, size),
		mutex: &sync.Mutex{},
	}
}

func (f *File) Combine(data []byte, index int) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.Data[index] = data
}

func (f *File) Save(path string, name string) error {
	if _, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}

	combined := []byte{}
	for i := 0; i < len(f.Data); i++ {
		combined = append(combined, f.Data[i]...)
	}

	return os.WriteFile(filepath.Join(path, name), combined, os.ModePerm)
}
