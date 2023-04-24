package storage

import (
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
)

type Storage struct {
	data []download.Download
	*setting.Settings
}

func New(s *setting.Settings) Storage {
	return Storage{
		data:     []download.Download{},
		Settings: s,
	}
}

func (s *Storage) Init() {
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

		_, err = os.Create(s.DataFilename)
		if err != nil {
			log.Fatalf("Cannot creating the file: %v", err)
		}
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

	err = json.Unmarshal(value, &s.data)
	if err != nil {
		log.Printf("Error opening data file: %v", err)
		return []download.Download{}
	}

	return s.data
}

func (s *Storage) Save(data []download.Download) {
	dataVal, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("Error marshaling the data: %v", err)
		return
	}

	err = os.WriteFile(s.DataFilename, dataVal, fs.ModePerm)
	if err != nil {
		log.Fatalf("Error writing the data into file: %v", err)
		return
	}
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
