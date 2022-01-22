package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golangplayground/app/core"
	"golangplayground/app/helpers"
	"golangplayground/app/models"
	"golangplayground/app/schemas"
	"net/http"
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

		db := app.Database.Find(&tasks, "todo_id = ?", todoId)

		paginator := helpers.NewPaginator(&tasks, db, request, helpers.PAGE_SIZE)

		pagination := schemas.ApiPaginationSchema{
			TotalItems: paginator.TotalItems,
			Next:       paginator.NextPage,
			Previous:   paginator.PreviousPage,
			Last:       paginator.LastPage,
			Current:    paginator.CurrentPage,
			Items:      paginator.Items,
		}

		serialized, _ := json.Marshal(pagination)

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}
