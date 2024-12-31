package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/daniyalumer/todo-list-go-chi/internal/api"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	"github.com/daniyalumer/todo-list-go-chi/internal/service"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} rq.Todo
// @Failure 400 {string} string "Bad Request"
// @Router /todo [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := service.ReadTodoList()
	if err != nil {
		log.Printf("error reading todo list: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.ParseResponse(w, todos, http.StatusOK)
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param todo body rq.Todo true "Todo"
// @Success 200 {string} string "Successfully created todo"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /todo/user/{user_id} [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var body rq.Todo
	err := api.ParseRequest(r, &body)
	if err != nil {
		log.Printf("error parsing request: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := api.ParseURLParameter(r, "user_id")
	if err != nil {
		log.Printf("error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting user_id to integer: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo, err := service.CreateTodo(uint(userID), body)
	if err != nil {
		log.Printf("error creating todo: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully created todo with id: %d", todo.ID), http.StatusOK)
}

// UpdateTodo godoc
// @Summary Update an existing todo
// @Description Update an existing todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo_id path int true "Todo ID"
// @Param todo body rq.TodoUpdate true "Todo Update"
// @Success 200 {string} string "Successfully updated todo"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /todo/{todo_id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var body rq.TodoUpdate
	err := api.ParseRequest(r, &body)
	if err != nil {
		log.Printf("error parsing request: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := api.ParseURLParameter(r, "todo_id")
	if err != nil {
		log.Printf("error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting todo_id to integer: %v", err)
		api.ParseResponse(w, fmt.Errorf("unable to process"), http.StatusInternalServerError)
		return
	}

	todo, err := service.UpdateTodo(uint(todoID), body)
	if err != nil {
		log.Printf("error updating todo: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully updated todo with id: %d", todo.ID), http.StatusOK)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo_id path int true "Todo ID"
// @Success 200 {string} string "Successfully deleted todo"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /todo/{todo_id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := api.ParseURLParameter(r, "todo_id")
	if err != nil {
		log.Printf("error parsing URL parameter: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting todo_id to integer: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo, err := service.DeleteTodo(uint(todoID))
	if err != nil {
		log.Printf("error deleting todo: %v", err)
		api.ParseResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.ParseResponse(w, fmt.Sprintf("successfully deleted todo with id: %d", todo.ID), http.StatusOK)
}
