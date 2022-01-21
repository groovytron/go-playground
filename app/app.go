package app

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	Database *gorm.DB
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	app.Router.StrictSlash(true)
	app.Router.HandleFunc("/todos", app.TodoIndexHandler()).Methods("GET")
	app.Router.HandleFunc("/todos/{todoId:[0-9]+}", app.TodoDetailsHandler()).Methods("GET")
	app.Router.HandleFunc("/todos/{todoId:[0-9]+}/tasks", app.TodoTasksHandler()).Methods("GET")
}
