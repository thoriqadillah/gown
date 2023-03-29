package fs

import (
	"os"
	"path/filepath"

	"github.com/thoriqadillah/gown/config"
	"github.com/thoriqadillah/gown/http"
)

type File struct {
	Data [][]byte
	*config.Config
}

func New(size int64, config *config.Config) *File {
	return &File{
		Data:   make([][]byte, size),
		Config: config,
	}
}

func (f *File) Combine(chunk *http.Chunk, index int) {
	f.Data[index] = chunk.Data
}

func (f *File) Save(name string) error {
	if _, err := os.Stat(f.SaveLocation); err != nil {
		if err := os.MkdirAll(f.SaveLocation, os.ModePerm); err != nil {
			return err
		}
	}

	combined := []byte{}
	for i := 0; i < len(f.Data); i++ {
		combined = append(combined, f.Data[i]...)
	}

	return os.WriteFile(filepath.Join(f.SaveLocation, name), combined, os.ModePerm)
}
