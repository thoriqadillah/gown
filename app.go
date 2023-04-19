package main

import (
	"changeme/gown/http"
	"changeme/gown/http/chunk"
	"changeme/gown/lib/factory/download"
	"changeme/gown/setting"
	"changeme/gown/storage"
	"changeme/gown/worker"
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
	"sync"
	"time"
)

// App struct
type App struct {
	ctx      context.Context
	settings setting.Settings
	pool     worker.Pool
	wg       sync.WaitGroup
	data     []download.Download
	err      error
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
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// creating the data folder
	if _, err := os.Stat(a.settings.DataLocation); err != nil {
		err := os.MkdirAll(a.settings.DataLocation, os.ModePerm)
		if err != nil {
			log.Fatalf("Cannot creating the folder: %v", err)
		}

		_, err = os.Create(a.settings.DataFilename)
		if err != nil {
			log.Fatalf("Cannot creating the file: %v", err)
		}
	}

	// reading the data file
	go func() {
		jsonFile, err := os.Open(a.settings.DataFilename)
		if err != nil {
			log.Printf("Error opening data file: %v", err)
			return
		}

		value, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Printf("Error opening data file: %v", err)
		}

		err = json.Unmarshal(value, &a.data)
		if err != nil {
			log.Printf("Error opening data file: %v", err)
		}
		log.Println("data[]", a.data)
	}()
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
	if len(a.data) == 0 {
		return []download.Download{}
	}

	return a.data
}

func (a *App) InitSetting() setting.Settings {
	return a.settings
}

func (a *App) Download(toDownload download.Download) error {
	start := time.Now()

	a.pool.Start()

	go func() {
		a.data = append(a.data, toDownload)
		dataVal, err := json.MarshalIndent(a.data, "", " ")
		if err != nil {
			log.Fatalf("Error marshaling the data: %v", err)
		}

		err = os.WriteFile(a.settings.DataFilename, dataVal, fs.ModePerm)
		if err != nil {
			log.Fatalf("Error writing the data into file: %v", err)
		}
	}()

	storage := storage.NewFile(toDownload.Metadata.Totalpart, &a.settings)
	chunks := make([]*chunk.Chunk, toDownload.Metadata.Totalpart)
	for part := range chunks {
		chunks[part] = chunk.New(a.ctx, toDownload, part, &a.settings, &a.wg)
	}

	for _, job := range chunks {
		a.wg.Add(1)
		a.pool.Add(job)
	}

	a.wg.Wait()

	for part, chunk := range chunks {
		storage.CombineFile(chunk.Data(), part)
	}

	if err := storage.SaveFile(toDownload.Name); err != nil {
		log.Printf("Error saving file: %v", err)
		return err
	}

	elapsed := time.Since(start)

	log.Printf("Took %v s to download %s\n", elapsed.Seconds(), toDownload.Name)

	return nil
}

/*
TODO:
- implement progress bar
- implement queue
- implement theming
- implement browser extension
- implement simultanous download
*/
