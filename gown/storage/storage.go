package storage

import (
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/goccy/go-json"
)

var settingBucket = []byte("settings")
var downloadBucket = []byte("downloads")
var themeBucket = []byte("themes")

type Storage struct {
	db *bolt.DB
}

func New() (*Storage, error) {
	datapath := fmt.Sprintf("%s/.gown/store.db", os.Getenv("HOME"))

	db, err := bolt.Open(datapath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(downloadBucket); err != nil {
			return fmt.Errorf("could not create downloads bucket: %v", err)
		}

		if _, err := tx.CreateBucketIfNotExists(settingBucket); err != nil {
			return fmt.Errorf("could not create settings bucket: %v", err)
		}

		if _, err := tx.CreateBucketIfNotExists(themeBucket); err != nil {
			return fmt.Errorf("could not create themes bucket: %v", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not setup bucket: %v", err)
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) CreateSetting() (stg setting.Settings) {
	err := s.db.View(func(tx *bolt.Tx) error {
		val := tx.Bucket(settingBucket).Get([]byte("setting"))

		if err := json.Unmarshal(val, &stg); err != nil {
			return fmt.Errorf("could not unmarshal setting: %v", err)
		}

		return nil
	})

	if err == nil {
		return stg
	}

	setting := setting.Default()
	s.db.Update(func(tx *bolt.Tx) error {
		val, err := json.Marshal(setting)
		if err != nil {
			return fmt.Errorf("could not marshaling setting: %v", err)
		}

		if err := tx.Bucket(settingBucket).Put([]byte("setting"), val); err != nil {
			return fmt.Errorf("could not store initial setting: %v", err)
		}

		return nil
	})

	return setting
}

func (s *Storage) Set(id string, data download.Download) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		val, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("could not store the download: %v", err)
		}

		return tx.Bucket(downloadBucket).Put([]byte(id), val)
	})
}

func (s *Storage) Delete(id string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(downloadBucket).Delete([]byte(id))
	})
}

func (s *Storage) GetAll() download.Store {
	all := download.Store{}
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(downloadBucket)
		c := bucket.Cursor()

		var download download.Download
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &download); err != nil {
				return fmt.Errorf("could not marshal the data: %v", err)
			}

			all[download.ID] = download
		}

		return nil
	})

	if err != nil {
		log.Println("no data found", err)
		return nil
	}

	return all
}
