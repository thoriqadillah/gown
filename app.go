package main

import (
	"changeme/gown/http"
	"changeme/gown/http/chunk"
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"changeme/gown/storage"
	"changeme/gown/worker"
	"context"
	"fmt"
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
	storage  *storage.Storage
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
	if _, err := os.Stat(a.settings.SaveLocation); err != nil {
		if err := os.MkdirAll(a.settings.SaveLocation, os.ModePerm); err != nil {
			log.Fatalf("Cannot creating the save location folder: %v", err)
		}
	}

	if _, err := os.Stat(a.settings.DataLocation); err != nil {
		err := os.MkdirAll(a.settings.DataLocation, os.ModePerm)
		if err != nil {
			log.Fatalf("Cannot creating the folder: %v", err)
		}
	}

	if _, err := os.OpenFile(a.settings.DataFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
		log.Fatalf("Cannot creating the file: %v", err)
	}

	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {

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

func (a *App) InitData() download.Store {
	return a.storage.Get()
}

func (a *App) Delete(name string) error {
	return a.storage.Delete(name)
}

func (a *App) UpdateData(data download.Store) {
	a.storage.Update(data)
}

func (a *App) InitSetting() setting.Settings {
	return a.settings
}

func (a *App) Download(toDownload *download.Download, resumepos []int64) error {
	canceled := false

	worker, err := worker.New(toDownload.Metadata.Totalpart, 1)
	if err != nil {
		log.Printf("Error creating worker: %v", err)
		return err
	}

	worker.Start()
	var wg sync.WaitGroup

	if err := a.storage.Add(toDownload.ID, *toDownload); err != nil {
		return err
	}

	chunks := make([]*chunk.Chunk, toDownload.Metadata.Totalpart)
	for part := range chunks {
		chunks[part] = chunk.New(a.ctx, toDownload, part, &a.settings, &wg)
		if len(resumepos) > 0 {
			chunks[part].ResumeFrom(resumepos[part])
		}
	}

	for _, job := range chunks {
		wg.Add(1)
		worker.Add(job)
	}

	runtime.EventsOn(a.ctx, "stop", func(optionalData ...interface{}) {
		canceled = true
	})

	wg.Wait()
	worker.Stop()

	if canceled {
		return nil
	}

	// combining
	runtime.EventsEmit(a.ctx, "downloaded", toDownload.ID, false)

	if err := storage.CreateFile(toDownload, &a.settings); err != nil {
		log.Printf("Error saving file: %v", err)
		return err
	}

	// combined
	runtime.EventsEmit(a.ctx, "downloaded", toDownload.ID, true)

	return nil
}

func (a *App) DeleteTempfile(toDelete download.Download) {
	for i := 0; i < toDelete.Metadata.Totalpart; i++ {
		if err := os.Remove(filepath.Join(a.settings.SaveLocation, fmt.Sprintf("%s-%d", toDelete.ID, i))); err != nil {
			log.Printf("Error deleting temp file of %s\n", toDelete.Name)
		}
	}
}

/*
TODO:
- implement queue
- implement theming
- implement browser extension
*/
