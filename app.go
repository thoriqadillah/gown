package main

import (
	"changeme/gown/http"
	"changeme/gown/setting"
	"changeme/gown/worker"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	setting.Settings
	worker.Pool
	err error
}

// NewApp creates a new App application struct
func NewApp() *App {
	s := setting.New()
	worker, err := worker.New(s.Concurrency, s.SimmultanousNum)

	return &App{
		Settings: s,
		Pool:     worker,
		err:      err,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Theme() setting.Themes {
	return a.Themes
}

func (a *App) Fetch(url string) (*http.Response, error) {
	res, err := http.Fetch(url, &a.Settings)
	if err != nil {
		return nil, err
	}

	return res, nil
}
