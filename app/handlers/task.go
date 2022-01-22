package handlers

import (
	"encoding/json"
	"fmt"
	"golangplayground/app/core"
	"golangplayground/app/models"
	"golangplayground/app/schemas"
	"net/http"
	"github.com/gorilla/mux"
)

func TodoTasksHandler(app *core.App) http.HandlerFunc {
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

		pagination := schemas.ApiPaginationSchema{
			Next:     nil,
			Previous: nil,
			Last:     1,
			Current:  1,
			Items:    tasks,
		}

		serialized, _ := json.Marshal(pagination)

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}
