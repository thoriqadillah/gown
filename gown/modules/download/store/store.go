package store

import "changeme/gown/modules/download"

type Store interface {
	GetAllData() download.Lists
	UpdateAllData(data download.Lists) error
}
