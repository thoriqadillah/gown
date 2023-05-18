package store

import "changeme/gown/modules/download"

type Store interface {
	GetAllData() download.Store
	UpdateAllData(data download.Store) error
}
