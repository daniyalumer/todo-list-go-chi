package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func TodoRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/get", GetTodosHandler)
	r.Post("/post", CreateTodoHandler)
	r.Put("/put", UpdateTodoHandler)
	r.Delete("/delete", DeleteTodoHandler)

	return r
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	message, todos := ReadTodoList()
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string `json:"message"`
		Todos   []Todo `json:"todos"`
	}{
		Message: message,
		Todos:   todos,
	}

	json.NewEncoder(w).Encode(response)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}

	todoDescription := r.FormValue("description")

	message, todo := CreateTodo(todoDescription)
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string `json:"message"`
		ToDo    Todo   `json:"todo"`
	}{
		Message: message,
		ToDo:    todo,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}

	todoDescription := r.FormValue("description")
	completedStr := r.FormValue("completed")
	Completed := completedStr == "true"

	IdStr := r.FormValue("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	message, todo := UpdateTodo(Id, Completed, todoDescription)
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string `json:"message"`
		ToDo    Todo   `json:"todo"`
	}{
		Message: message,
		ToDo:    todo,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}

	IdStr := r.FormValue("id")
	
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	message, todo := DeleteTodo(Id)
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string `json:"message"`
		ToDo    Todo   `json:"todo"`
	}{
		Message: message,
		ToDo:    todo,
	}
	json.NewEncoder(w).Encode(response)
}
