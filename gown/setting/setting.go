package setting

const (
	mb                                   = 1024 * 1024
	DEFAULT_PART_SIZE              int64 = 5 * mb // 5 MB
	DEFAULT_CONCURRENCY                  = 100
	DEFAULT_MAX_TRIES                    = 3
	DEFAULT_SIMMULATANOUS_DOWNLOAD       = 4
	DEFAULT_SAVE_LOCATION                = "~/Downloads/"
)

type Settings struct {
	Themes
	Partsize        int64  `json:"partsize"`
	Concurrency     int    `json:"concurrency"`
	Maxtries        int    `json:"maxtries"`
	SimmultanousNum int    `json:"simmultanousNum"`
	SaveLocation    string `json:"saveLocation"`
}

func New() Settings {
	return Settings{
		Themes:          Theme(),
		Partsize:        DEFAULT_PART_SIZE,
		Concurrency:     DEFAULT_CONCURRENCY,
		Maxtries:        DEFAULT_MAX_TRIES,
		SimmultanousNum: DEFAULT_SIMMULATANOUS_DOWNLOAD,
		SaveLocation:    DEFAULT_SAVE_LOCATION,
	}
}
