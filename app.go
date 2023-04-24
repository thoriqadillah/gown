package main

import (
	"changeme/gown/http"
	"changeme/gown/http/chunk"
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"changeme/gown/storage"
	"changeme/gown/worker"
	"context"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	settings setting.Settings
	storage  storage.Storage
}

// NewApp creates a new App application struct
func NewApp() *App {
	s := setting.New()

	return &App{
		settings: s,
		storage:  storage.New(&s),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.storage.Init()
}

func (a *App) Theme() setting.Theme {
	return a.settings.Themes
}

func (a *App) Fetch(url string) (*download.Download, error) {
	res, err := http.Fetch(url, &a.settings)
	if err != nil {
		return nil, err
	}

	factory := download.NewFactory(res)
	data := factory.Create()

	return &data, nil
}

func (a *App) InitData() []download.Download {
	return a.storage.Get()
}

func (a *App) UpdateName(oldname string, newname string) error {
	oldname = filepath.Join(a.settings.SaveLocation, oldname)
	newname = filepath.Join(a.settings.SaveLocation, newname)

	if _, err := os.Stat(oldname); err != nil {
		return err
	}

	log.Println(oldname, newname)
	if err := os.Rename(oldname, newname); err != nil {
		return err
	}

	return nil
}

func (a *App) UpdateData(data []download.Download) {
	a.storage.Save(data)
}

func (a *App) InitSetting() setting.Settings {
	return a.settings
}

func (a *App) Download(toDownload *download.Download) error {
	worker, err := worker.New(toDownload.Metadata.Totalpart, a.settings.SimmultanousNum)
	if err != nil {
		log.Printf("Error initializing worker: %v\n", err)
	}

	worker.Start()

	var wg sync.WaitGroup

	data := a.storage.Get()
	data = append([]download.Download{*toDownload}, data...)
	a.storage.Save(data)

	storage := storage.NewFile(toDownload.Metadata.Totalpart, &a.settings)
	chunks := make([]*chunk.Chunk, toDownload.Metadata.Totalpart)
	for part := range chunks {
		chunks[part] = chunk.New(a.ctx, *toDownload, part, &a.settings, &wg)
	}

	for _, job := range chunks {
		wg.Add(1)
		worker.Add(job)
	}

	wg.Wait()
	worker.Stop()

	// combining
	runtime.EventsEmit(a.ctx, "downloaded", toDownload.ID, false)

	for part, chunk := range chunks {
		storage.CombineFile(chunk.Data(), part)
	}

	if err := storage.SaveFile(toDownload.Name); err != nil {
		log.Printf("Error saving file: %v", err)
		return err
	}

	// combined
	runtime.EventsEmit(a.ctx, "downloaded", toDownload.ID, true)
	runtime.EventsOff(a.ctx, "downloaded")

	return nil
}

/*
TODO:
- implement queue
- implement theming
- implement browser extension
*/
