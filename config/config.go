package config

const (
	DEFAULT_PART_SIZE              = 1024 * 1024 * 5 // 5 MB
	DEFAULT_CONCURRENCY            = 10
	DEFAULT_MAX_TRIES              = 3
	DEFAULT_SIMMULATANOUS_DOWNLOAD = 1
	DEFAULT_SAVE_LOCATION          = "./storage"
)

type Config struct {
	Partsize        int64
	Concurrency     int
	Maxtries        int
	SimmultanousNum int
	SaveLocation    string
}

func Default() *Config {
	return &Config{
		Partsize:        DEFAULT_PART_SIZE,
		Concurrency:     DEFAULT_CONCURRENCY,
		Maxtries:        DEFAULT_CONCURRENCY,
		SimmultanousNum: DEFAULT_SIMMULATANOUS_DOWNLOAD,
		SaveLocation:    DEFAULT_SAVE_LOCATION,
	}
}
