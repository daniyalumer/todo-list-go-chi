package service

import (
	"fmt"
	"time"

	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	repository "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateTodo(userID uint, body rq.TodoCreate) (dao.Todo, error) {
	var user dao.User

	if err := repository.CheckDeleted(&user, userID); err != nil {
		return dao.Todo{}, fmt.Errorf("failed to find user: %v", err)
	}

	newTodo := dao.Todo{
		Description: body.Description,
		CompletedAt: nil,
		Completed:   false,
		UserID:      userID,
	}
	err := repository.Create(&newTodo)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to create todo: %v", err)
	}
	return newTodo, nil
}

func ReadTodoList() ([]dao.Todo, error) {
	var todos []dao.Todo
	err := repository.FindAll(&todos)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %v", err)
	}
	return todos, nil
}

func UpdateTodo(todoID uint, body rq.TodoUpdate) (dao.Todo, error) {
	var todo dao.Todo
	if err := repository.FindById(&todo, todoID); err != nil {
		return dao.Todo{}, fmt.Errorf("todo not found: %v", err)
	}
	updates := map[string]interface{}{}
	if body.Description != "" {
		updates["description"] = body.Description
	}

	if body.Completed {
		updates["completed"] = body.Completed
		updates["completed_at"] = time.Now()
	}

	if err := repository.Update(&todo, updates); err != nil {
		return dao.Todo{}, fmt.Errorf("failed to update todo: %v", err)
	}

	return todo, nil
}

func DeleteTodo(todoID uint) (dao.Todo, error) {
	var todo dao.Todo

	err := repository.FindById(&todo, todoID)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to find todo: %v", err)
	}

	err = repository.Delete(&todo)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to delete todo: %v", err)
	}

	return todo, nil
}
