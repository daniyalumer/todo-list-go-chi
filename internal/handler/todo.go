package handler

import (
	"net/http"

	"github.com/daniyalumer/todo-list-go-chi/internal/helper"
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
		helper.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	helper.ResponseWriter(w, todos, http.StatusOK)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	todoDescription := r.Form.Get("description")

	useridstr := r.URL.Query().Get("userId")

	userID, err := helper.ConvertIdToInteger(useridstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.CreateTodo(todoDescription, userID)
	if err != nil {
		helper.ResponseWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.ResponseWriter(w, todo, http.StatusOK)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	todoDescription := r.Form.Get("description")
	completedStr := r.Form.Get("completed")

	Completed := completedStr == "true"

	idstr := r.URL.Query().Get("id")
	Id, err := helper.ConvertIdToInteger(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.UpdateTodo(Id, Completed, todoDescription)
	if err != nil {
		helper.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	helper.ResponseWriter(w, todo, http.StatusOK)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if !helper.ParseForm(w, r) {
		return
	}

	idstr := r.URL.Query().Get("id")

	Id, err := helper.ConvertIdToInteger(idstr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := service.DeleteTodo(Id)
	if err != nil {
		helper.ResponseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}
	helper.ResponseWriter(w, todo, http.StatusOK)
}
