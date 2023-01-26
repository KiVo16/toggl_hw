package app

import "base/internal/app/handlers"

type App struct {
	Handlers handlers.Handlers
}

func NewApp(handlers handlers.Handlers) *App {

	return &App{
		Handlers: handlers,
	}
}
