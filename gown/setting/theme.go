package setting

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const DEFAULT_THEME_PATH = "/home/thoriqadillah/Development/Go/gown/themes/default.json"

type Theme struct {
	Component
}

func Themes() Theme {
	return Theme{
		Component: theme(DEFAULT_THEME_PATH),
	}
}

func theme(path string) Component {
	f, err := os.Open(path)
	if err != nil {
		log.Panicf("Error opening file: %v\n", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Panicf("Error reading file: %v\n", err)
	}

	var components Component
	if err := json.Unmarshal(data, &components); err != nil {
		log.Panicf("Error unmarshalling file: %v\n", err)
	}

	return components
}

type Component struct {
	TextColor       string `json:"textColor"`
	BackgroundColor string `json:"backgroundColor"`
}

// TODO: store theme into disk
// TODO: make the theme usable
