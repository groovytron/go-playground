package app

import (
	"github.com/gorilla/mux"
	"golangplayground/app/core"
	"golangplayground/app/handlers"
)

func NewApp() *core.App {
	app := &core.App{
		Router: mux.NewRouter(),
	}

	initAppRoutes(app)

	return app
}

func initAppRoutes(app *core.App) {
	app.Router.StrictSlash(true)
	app.Router.HandleFunc("/api/todos", handlers.TodoIndexHandler(app)).Methods("GET")
	app.Router.HandleFunc("/api/todos", handlers.TodoSaveHandler(app)).Methods("POST")
	app.Router.HandleFunc("/api/todos/{todoId:[0-9]+}", handlers.TodoDetailsHandler(app)).Methods("GET")
	app.Router.HandleFunc("/api/todos/{todoId:[0-9]+}/tasks", handlers.TodoTasksHandler(app)).Methods("GET")
}
