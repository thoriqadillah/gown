package setting

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const DEFAULT_THEME_PATH = "/home/thoriqadillah/Development/Go/gown/themes/default.json"

type Themes struct {
	Components
}

func Theme() Themes {
	return Themes{
		Components: theme(DEFAULT_THEME_PATH),
	}
}

func theme(path string) Components {
	f, err := os.Open(path)
	if err != nil {
		log.Panicf("Error opening file: %v\n", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Panicf("Error reading file: %v\n", err)
	}

	var components Components
	if err := json.Unmarshal(data, &components); err != nil {
		log.Panicf("Error unmarshalling file: %v\n", err)
	}

	return components
}

type Components struct {
	TextColor       string `json:"textColor"`
	BackgroundColor string `json:"backgroundColor"`
}
