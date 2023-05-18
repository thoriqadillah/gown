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

	filename := fmt.Sprintf("%s/setting.json", init.DataLocation)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Could not create or open the setting file: %v", err)
	}

	if err := json.NewDecoder(file).Decode(&init); err != nil {
		log.Printf("Could not decode setting file: %v", err)
	}

	return init
}

func (s *fileStore) UpdateSetting(data *setting.Settings) error {
	filename := fmt.Sprintf("%s/setting.json", data.DataLocation)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open the setting file: %v", err)
	}

	return json.NewEncoder(file).Encode(data)
}

func (s *fileStore) DefaultSetting() *setting.Settings {
	return setting.Default()
}
