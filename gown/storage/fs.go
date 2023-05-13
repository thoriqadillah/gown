package storage

import (
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateFile(toSave *download.Download, setting *setting.Settings) error {
	file, err := os.OpenFile(filepath.Join(setting.SaveLocation, toSave.Name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for i := 0; i < toSave.Metadata.Totalpart; i++ {
		filename := filepath.Join(setting.SaveLocation, fmt.Sprintf("%s-%d", toSave.ID, i))
		tmpFile, err := os.Open(filename)
		if err != nil {
			return err
		}

		if _, err := io.Copy(file, tmpFile); err != nil {
			return err
		}

		if err := os.Remove(filename); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) DeleteFile(filename string) error {
	if err := os.Remove(filename); err != nil {
		return err
	}

	return nil
}
