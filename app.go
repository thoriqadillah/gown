package main

import (
	factory "changeme/gown/lib/factory/download"
	"changeme/gown/lib/fs"
	"changeme/gown/lib/worker"
	"changeme/gown/modules/download"
	"changeme/gown/modules/download/chunk"
	"changeme/gown/modules/setting"
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
	settings *setting.Settings
	worker   worker.Pool
}

// NewApp creates a new App application struct
func NewApp(s *setting.Settings) *App {
	worker, err := worker.New(s.Concurrency, s.SimmultanousNum)
	if err != nil {
		log.Printf("Error creating worker: %v", err)
	}

	return &App{
		settings: s,
		worker:   worker,
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

	a.ctx = ctx
	a.worker.Start()
}

func (a *App) shutdown(ctx context.Context) {
	a.worker.Stop()
}

func (a *App) Fetch(url string) (*download.Download, error) {
	res, err := download.Fetch(url, a.settings)
	if err != nil {
		return nil, err
	}

	factory := factory.NewFactory(res)
	data := factory.Create()

	return &data, nil
}

func (a *App) Download(toDownload *download.Download, resume bool) error {
	canceled := false

	var wg sync.WaitGroup

	chunks := make([]*chunk.Chunk, toDownload.Metadata.Totalpart)
	for part := range chunks {
		chunks[part] = chunk.New(a.ctx, toDownload, part, a.settings, &wg)

		if resume {
			//TODO: recheck if the resumed url is broken or not by comparing the size of the original download and newly fetched data
			tempFile := filepath.Join(a.settings.SaveLocation, fmt.Sprintf("%s-%d", toDownload.ID, part))
			f, err := os.Stat(tempFile)
			if err != nil {
				return err
			}

			chunks[part].ResumeFrom(f.Size())
		}
	}

	for _, job := range chunks {
		wg.Add(1)
		a.worker.Add(job)
	}

	runtime.EventsOn(a.ctx, "stop", func(optionalData ...interface{}) {
		canceled = true
	})

	wg.Wait()

	if canceled {
		return nil
	}

	// combining
	runtime.EventsEmit(a.ctx, "downloaded", toDownload.ID, false)

	if err := fs.CreateFile(toDownload, a.settings); err != nil {
		log.Printf("Error saving file: %v", err)
		return err
	}

	// combined
	runtime.EventsEmit(a.ctx, "downloaded", toDownload.ID, true)

	return nil
}

/*
TODO:
- implement queue
- implement theming
- implement browser extension
*/
