package handlers

import (
	"encoding/json"
	"fmt"
	"golangplayground/app/core"
	"golangplayground/app/helpers"
	"golangplayground/app/models"
	"golangplayground/app/schemas"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func TodoIndexHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var todos []models.Todo

		db := app.Database.Find(&todos).Order("id ASC")

		paginator := helpers.NewPaginator(&todos, db, request, helpers.PAGE_SIZE)

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

func TodoDetailsHandler(app *core.App) http.HandlerFunc {
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

func TodoCreateHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var todo schemas.TodoCreateSchema

		err := json.NewDecoder(request.Body).Decode(&todo)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)

			return
		}

		validator := validator.New()

		err = validator.Struct(todo)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		todoToSave := models.Todo{Name: todo.Name, Description: todo.Description}

		app.Database.Create(&todoToSave)

		serialized, _ := json.Marshal(todoToSave)

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}

func TodoUpdateHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		todoId := routeParams["todoId"]

		var todoItem models.Todo

		result := app.Database.Find(&todoItem, todoId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		var todo schemas.TodoCreateSchema

		err := json.NewDecoder(request.Body).Decode(&todo)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		validator := validator.New()

		err = validator.Struct(todo)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		todoItem.Name = todo.Name
		todoItem.Description = todo.Description

		app.Database.Save(&todoItem)

		serialized, _ := json.Marshal(todoItem)

		writer.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(writer, string(serialized))
	}
}

func TodoDeleteHandler(app *core.App) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		routeParams := mux.Vars(request)

		todoId := routeParams["todoId"]

		var todoItem models.Todo

		result := app.Database.Find(&todoItem, todoId)

		if result.RowsAffected == 0 {
			writer.WriteHeader(http.StatusNotFound)

			return
		}

		app.Database.Select("Tasks").Unscoped().Delete(&todoItem)
	}
}
