package storage

import (
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/goccy/go-json"
)

var instance *Storage
var mutex = &sync.Mutex{}

type Storage struct {
	*setting.Settings
}

func New(s *setting.Settings) *Storage {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		instance = &Storage{
			Settings: s,
		}
	}

	return instance
}

func (s *Storage) Get() download.Store {
	jsonFile, err := os.Open(s.DataFilename)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return download.Store{}
	}

	var store download.Store
	if err := json.NewDecoder(jsonFile).Decode(&store); err != nil {
		return download.Store{}
	}

	return store
}

func (s *Storage) Add(id string, val download.Download) error {
	data := s.Get()
	data[id] = val

	return s.Update(data)
}

func (s *Storage) Update(data download.Store) error {
	jsonFile, err := os.OpenFile(s.DataFilename, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return err
	}

	return json.NewEncoder(jsonFile).Encode(data)
}

func (s *Storage) Delete(name string) error {
	filename := filepath.Join(s.SaveLocation, name)
	if err := os.Remove(filename); err != nil {
		return err
	}

	return nil
}

// TODO: implement storing data into persistent file
// The data will be stored in a JSON file
// The file will be stored in the data directory
// The data will be :
// - list of downloaded file
// - list of queued download
// - list of failed download
// - selected theme
// - theme configuration
// - TBD
