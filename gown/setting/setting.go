package setting

const (
	mb                                   = 1024 * 1024
	DEFAULT_PART_SIZE              int64 = 5 * mb // 5 MB
	DEFAULT_CONCURRENCY                  = 100
	DEFAULT_MAX_TRIES                    = 3
	DEFAULT_SIMMULATANOUS_DOWNLOAD       = 4
	DEFAULT_SAVE_LOCATION                = "/home/thoriqadillah/Downloads/"
	DEFAULT_DATA_LOCATION                = "/home/thoriqadillah/.gown/"
	DEFAULT_DATA_FILE_NAME               = "/home/thoriqadillah/.gown/data.json"
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
}

func New() Settings {
	return Settings{
		Themes:          Themes(),
		Partsize:        DEFAULT_PART_SIZE,
		Concurrency:     DEFAULT_CONCURRENCY,
		Maxtries:        DEFAULT_MAX_TRIES,
		SimmultanousNum: DEFAULT_SIMMULATANOUS_DOWNLOAD,
		SaveLocation:    DEFAULT_SAVE_LOCATION,
		DataLocation:    DEFAULT_DATA_LOCATION,
		DataFilename:    DEFAULT_DATA_FILE_NAME,
	}
}
