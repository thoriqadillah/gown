package storage

import (
	"changeme/gown/setting"
	"os"
	"path/filepath"
)

type File struct {
	data    [][]byte
	setting *setting.Settings
}

func NewFile(parts int, setting *setting.Settings) File {
	return File{
		data:    make([][]byte, parts),
		setting: setting,
	}
}

func (s *File) CombineFile(data []byte, index int) {
	s.data[index] = data
}

func (s *File) SaveFile(name string) error {
	if _, err := os.Stat(s.setting.SaveLocation); err != nil {
		if err := os.MkdirAll(s.setting.SaveLocation, os.ModePerm); err != nil {
			return err
		}
	}

	combined := []byte{}
	for i := 0; i < len(s.data); i++ {
		combined = append(combined, s.data[i]...)
	}

	return os.WriteFile(filepath.Join(s.setting.SaveLocation, name), combined, os.ModePerm)
}
