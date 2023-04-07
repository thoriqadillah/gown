package main

import (
	"changeme/gown/setting"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	setting.Settings
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Settings: setting.New(),
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
