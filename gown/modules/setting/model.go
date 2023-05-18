package setting

import (
	"fmt"
	"os"
)

const (
	mb                                   = 1024 * 1024
	DEFAULT_PART_SIZE              int64 = 5 * mb // 5 MB
	DEFAULT_CONCURRENCY                  = 100
	DEFAULT_MAX_TRIES                    = 3
	DEFAULT_SIMMULATANOUS_DOWNLOAD       = 4
)

type Settings struct {
	Themes          Theme  `json:"themes"`
	Partsize        int64  `json:"partsize"`
	Concurrency     int    `json:"concurrency"`
	Maxtries        int    `json:"maxtries"`
	SimmultanousNum int    `json:"simmultanousNum"`
	SaveLocation    string `json:"saveLocation"`
	DataLocation    string `json:"dataLocation"`
	DataFilename    string `json:"dataFilename"`
	SettingFilename string `json:"settingFilename"`
}

func Default() *Settings {
	return &Settings{
		Themes:          Themes(),
		Partsize:        DEFAULT_PART_SIZE,
		Concurrency:     DEFAULT_CONCURRENCY,
		Maxtries:        DEFAULT_MAX_TRIES,
		SimmultanousNum: DEFAULT_SIMMULATANOUS_DOWNLOAD,
		SaveLocation:    fmt.Sprintf("%s/Downloads/", os.Getenv("HOME")),
		DataLocation:    fmt.Sprintf("%s/.gown/", os.Getenv("HOME")),
		DataFilename:    fmt.Sprintf("%s/.gown/data.json", os.Getenv("HOME")),
		SettingFilename: fmt.Sprintf("%s/.gown/setting.json", os.Getenv("HOME")),
	}
}

// TODO: store settings into disk
