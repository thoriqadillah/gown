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
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	settings setting.Settings
	pool     worker.Pool
	wg       sync.WaitGroup
	err      error
	storage  storage.Storage
}

// NewApp creates a new App application struct
func NewApp() *App {
	s := setting.New()
	worker, err := worker.New(s.Concurrency, s.SimmultanousNum)

	return &App{
		settings: s,
		pool:     worker,
		err:      err,
		wg:       sync.WaitGroup{},
		storage:  storage.New(&s),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.storage.Init()
}

func (a *App) shutdown(ctx context.Context) {
	a.pool.Stop()
}

func (a *App) Theme() setting.Themes {
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

func (a *App) UpdateData(data []download.Download) {
	a.storage.Save(data)
}

func (a *App) InitSetting() setting.Settings {
	return a.settings
}

func (a *App) Download(toDownload *download.Download) error {
	a.pool.Start()

	toDownload.Date = time.Now()
	data := a.storage.Get()
	data = append([]download.Download{*toDownload}, data...)
	a.storage.Save(data)

	storage := storage.NewFile(toDownload.Metadata.Totalpart, &a.settings)
	chunks := make([]*chunk.Chunk, toDownload.Metadata.Totalpart)
	for part := range chunks {
		chunks[part] = chunk.New(a.ctx, *toDownload, part, &a.settings, &a.wg)
	}

	for _, job := range chunks {
		a.wg.Add(1)
		a.pool.Add(job)
	}

	a.wg.Wait()

	runtime.EventsEmit(a.ctx, "done", toDownload.ID)

	for part, chunk := range chunks {
		storage.CombineFile(chunk.Data(), part)
	}

	if err := storage.SaveFile(toDownload.Name); err != nil {
		log.Printf("Error saving file: %v", err)
		return err
	}

	return nil
}

/*
TODO:
- implement queue
- implement theming
- implement browser extension
- implement simultanous download
*/
