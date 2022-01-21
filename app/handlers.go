package app

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"golangplayground/app/models"
	
)

type ApiPagination struct {
	Next     *int64      `json:"next"`
	Previous *int64      `json:"previous"`
	Last     int64       `json:"last"`
	Current  int64       `json:"current"`
	Items    interface{} `json:"items"`
}

func (app *App) TodoIndexHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var todosList []models.Todo

		app.Database.Find(&todosList)

		pagination := ApiPagination{
			Next: nil,
			Previous: nil,
			Last: 1,
			Current: 1,
			Items: todosList,
		}

		serialized, _ := json.Marshal(pagination)
		
		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}

func (app *App) TodoDetailsHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		todoId := routeParams["todoId"]

		var todoItem models.Todo

		result := app.Database.Find(&todoItem, todoId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		serialized, _ := json.Marshal(todoItem) 

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}


func (app *App) TodoTasksHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		var tasks []models.Task

		todoId := routeParams["todoId"]

		var todoItem models.Todo

		result := app.Database.Find(&todoItem, todoId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		app.Database.Find(&tasks, "todo_id = ?", todoId)

		serialized, _ := json.Marshal(tasks) 

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}
