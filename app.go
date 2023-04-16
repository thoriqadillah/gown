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
	"sync"
	"time"
)

// App struct
type App struct {
	ctx      context.Context
	settings setting.Settings
	pool     worker.Pool
	wg       sync.WaitGroup
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
}

func (a *App) shutdown(ctx context.Context) {
	a.pool.Stop()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Theme() setting.Themes {
	return a.settings.Themes
}

func (a *App) Fetch(url string) (*http.Response, error) {
	res, err := http.Fetch(url, &a.settings)
	if err != nil {
		return nil, err
	}

	// TODO: utilize this
	factory := download.NewFactory(res)
	log.Println(factory.Create())

	return res, nil
}

func (a *App) Download(res http.Response) error {
	start := time.Now()

	a.pool.Start()
	// defer a.pool.Stop()

	storage := storage.NewFile(res.Totalpart, &a.settings)
	chunks := make([]*chunk.Chunk, res.Totalpart)
	for part := range chunks {
		chunks[part] = chunk.New(a.ctx, res, part, &a.wg)
	}

	for _, job := range chunks {
		a.wg.Add(1)
		a.pool.Add(job)
	}

	a.wg.Wait()

	for part, chunk := range chunks {
		storage.CombineFile(chunk.Data(), part)
	}

	if err := storage.SaveFile(res.Filename); err != nil {
		log.Printf("Error saving file: %v", err)
		return err
	}

	elapsed := time.Since(start)
	log.Printf("Took %v s to download %s\n", elapsed.Seconds(), res.Filename)

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
