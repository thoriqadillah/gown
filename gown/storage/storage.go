package storage

import "changeme/gown/lib/factory/download"

type Storage struct {
	Queue     []download.Download `json:"queue"`
	Downloads []download.Download `json:"downloads"`
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
