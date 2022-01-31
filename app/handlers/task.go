package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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

		db := app.Database.Find(&tasks, "todo_id = ?", todoId).Order("id ASC")

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

func TaskCreateHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		todoId := routeParams["todoId"]

		var todoItem models.Todo

		result := app.Database.Find(&todoItem, todoId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		var task schemas.TaskCreateSchema

		err := json.NewDecoder(request.Body).Decode(&task)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)

			return
		}

		validator := validator.New()

		err = validator.Struct(task)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		taskToSave := models.Task{Name: task.Name, Description: task.Description, Todo: todoItem}

		app.Database.Create(&taskToSave)

		serialized, _ := json.Marshal(taskToSave)

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}

func TaskUpdateHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		taskId := routeParams["taskId"]

		var taskItem models.Task

		result := app.Database.Find(&taskItem, taskId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		var task schemas.TaskCreateSchema

		err := json.NewDecoder(request.Body).Decode(&task)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)

			return
		}

		validator := validator.New()

		err = validator.Struct(task)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		taskItem.Name = task.Name
		taskItem.Description = task.Description

		app.Database.Save(&taskItem)

		serialized, _ := json.Marshal(taskItem)

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}

func TaskDeleteHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		taskId := routeParams["taskId"]

		var task models.Task

		result := app.Database.Find(&task, taskId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		app.Database.Unscoped().Delete(&task)
	}
}
