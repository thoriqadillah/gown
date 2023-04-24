package storage

import (
	"changeme/gown/setting"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	tmpFile []string
	size    int64
	setting *setting.Settings
}

func NewFile(parts int, size int64, setting *setting.Settings) File {
	return File{
		tmpFile: make([]string, parts),
		size:    size,
		setting: setting,
	}
}

func (s *File) CombineFile(id string, index int) {
	s.tmpFile[index] = id
}

func (s *File) SaveFile(name string) error {
	// create the file for the downloaded data
	file, err := os.OpenFile(filepath.Join(s.setting.SaveLocation, name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for _, name := range s.tmpFile {
		filename := filepath.Join(s.setting.SaveLocation, name)
		tmpFile, err := os.Open(filename)
		if err != nil {
			return err
		}

		data, err := io.ReadAll(tmpFile)
		if err != nil {
			return err
		}

		if _, err := file.Write(data); err != nil {
			return err
		}

		if err := os.Remove(filename); err != nil {
			return err
		}
	}

	return nil
}
