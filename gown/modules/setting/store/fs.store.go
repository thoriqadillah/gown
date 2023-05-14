package store

import (
	"changeme/gown/modules/setting"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
)

type fileStore struct{}

func NewFileStore() Store {
	return &fileStore{}
}

func (s *fileStore) GetSetting() *setting.Settings {
	init := setting.Default()

	file, err := os.OpenFile(init.DataFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Could not create or open the setting file: %v", err)
	}

	if err := json.NewDecoder(file).Decode(&init); err != nil {
		log.Printf("Could not decode setting file: %v", err)
	}

	return init
}

func (s *fileStore) UpdateSetting(data *setting.Settings) error {
	file, err := os.OpenFile(data.DataFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("could not open the setting file: %v", err)
	}

	return json.NewEncoder(file).Encode(data)
}
