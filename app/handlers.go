package app

import (
	"encoding/json"
	"fmt"
	"golangplayground/app/models"
	"golangplayground/app/schemas"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


func (app *App) TodoIndexHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var todosList []models.Todo

		app.Database.Find(&todosList)

		pagination := schemas.ApiPaginationSchema{
			Next:     nil,
			Previous: nil,
			Last:     1,
			Current:  1,
			Items:    todosList,
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


func (app *App) TodoSaveHandler() http.HandlerFunc {
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

		fmt.Fprintf(writer, string(serialized))
	}
}
