package main

import (
	downloadstore "changeme/gown/modules/download/store"
	settingstore "changeme/gown/modules/setting/store"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	setting := settingstore.NewFileStore()
	s := setting.GetSetting()

	app := NewApp(s)

	download := downloadstore.NewFileStore(s)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Gown",
		MinWidth:  815,
		MinHeight: 590,
		Width:     815,
		Height:    590,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			setting,
			download,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
