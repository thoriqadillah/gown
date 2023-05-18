package store

import (
	"changeme/gown/modules/download"
	"changeme/gown/modules/setting"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

type fileStore struct {
	s *setting.Settings
}

func NewFileStore(s *setting.Settings) Store {
	if _, err := os.OpenFile(s.DataFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
		log.Fatalf("Cannot creating the file: %v", err)
	}

	return &fileStore{
		s: s,
	}
}

func (s *fileStore) GetAllData() (store download.Store) {
	jsonFile, err := os.Open(s.s.DataFilename)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return download.Store{}
	}

	if err := json.NewDecoder(jsonFile).Decode(&store); err != nil {
		return download.Store{}
	}

	return store
}

func (s *fileStore) UpdateAllData(data download.Store) error {
	jsonFile, err := os.OpenFile(s.s.DataFilename, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return err
	}

	return json.NewEncoder(jsonFile).Encode(data)
}

func (s *fileStore) DeleteFile(target download.Download) error {
	filename := filepath.Join(s.s.SaveLocation, target.Name)
	return os.Remove(filename)
}

func (s *fileStore) DeleteTempfile(target download.Download) {
	for i := 0; i < target.Metadata.Totalpart; i++ {
		if err := os.Remove(filepath.Join(s.s.SaveLocation, fmt.Sprintf("%s-%d", target.ID, i))); err != nil {
			log.Printf("Error deleting temp file of %s\n", target.Name)
		}
	}
}
