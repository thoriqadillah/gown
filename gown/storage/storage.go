package storage

import (
	"changeme/gown/setting"
	"os"
	"path/filepath"
)

type Storage struct {
	data    [][]byte
	setting *setting.Settings
}

func New(parts int, setting *setting.Settings) Storage {
	return Storage{
		data:    make([][]byte, parts),
		setting: setting,
	}
}

func (s *Storage) Combine(data []byte, index int) {
	s.data[index] = data
}

func (s *Storage) Save(name string) error {
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
