package storage

import (
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Storage struct {
	*setting.Settings
}

func New(s *setting.Settings) Storage {
	if _, err := os.Stat(s.SaveLocation); err != nil {
		if err := os.MkdirAll(s.SaveLocation, os.ModePerm); err != nil {
			log.Fatalf("Cannot creating the save location folder: %v", err)
		}
	}

	if _, err := os.Stat(s.DataLocation); err != nil {
		err := os.MkdirAll(s.DataLocation, os.ModePerm)
		if err != nil {
			log.Fatalf("Cannot creating the folder: %v", err)
		}
	}

	if _, err := os.OpenFile(s.DataFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
		log.Fatalf("Cannot creating the file: %v", err)
	}

	return Storage{
		Settings: s,
	}
}

func (s *Storage) Get() []download.Download {
	jsonFile, err := os.Open(s.DataFilename)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return []download.Download{}
	}

	value, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return []download.Download{}
	}

	var downloaded []download.Download
	err = json.Unmarshal(value, &downloaded)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return []download.Download{}
	}

	return downloaded
}

func (s *Storage) Add(val download.Download) error {
	data := s.Get()
	data = append([]download.Download{val}, data...)

	return s.Update(data)
}

func (s *Storage) Update(data []download.Download) error {
	dataVal, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Printf("Error marshaling the data: %v", err)
		return err
	}

	err = os.WriteFile(s.DataFilename, dataVal, fs.ModePerm)
	if err != nil {
		log.Printf("Error writing the data into file: %v", err)
		return err
	}

	return nil
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
