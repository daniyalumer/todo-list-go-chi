package controller

import (
	"net/http"
	"strconv"

	"github.com/daniyalumer/todo-list-go-chi/internal/helper"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
	"github.com/go-chi/chi/v5"
)

func TodoRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", GetTodosHandler)
	r.Post("/", CreateTodoHandler)
	r.Put("/", UpdateTodoHandler)
	r.Delete("/", DeleteTodoHandler)

	return r
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ReadTodoList()
	if err != nil {
		helper.TodoListResponseWriter(w, "failed to read todo list", []models.Todo{}, http.StatusBadRequest)
		return
	}
	helper.TodoListResponseWriter(w, "todo list read successfully", todos, http.StatusAccepted)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	todoDescription := r.FormValue("description")
	userIDStr := r.URL.Query().Get("userId")


	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.CreateTodo(todoDescription, userID)
	if err != nil {
		helper.TodoResponseWriter(w, "failed to create todo item", models.Todo{}, http.StatusBadRequest)
		return
	}
	helper.TodoResponseWriter(w, "created todo item successfully", todo, http.StatusAccepted)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	todoDescription := r.FormValue("description")
	completedStr := r.FormValue("completed")
	Completed := completedStr == "true"

	IdStr := r.URL.Query().Get("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.UpdateTodo(Id, Completed, todoDescription)
	if err != nil {
		helper.TodoResponseWriter(w, "failed to update todo item", models.Todo{}, http.StatusBadRequest)
		return
	}
	helper.TodoResponseWriter(w, "updated todo item successfully", todo, http.StatusAccepted)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	IdStr := r.URL.Query().Get("id")

	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.DeleteTodo(Id)
	if err != nil {
		helper.TodoResponseWriter(w, "failed to deleted todo item", models.Todo{}, http.StatusBadRequest)
		return
	}
	helper.TodoResponseWriter(w, "successfully deleted todo item", todo, http.StatusAccepted)
}
