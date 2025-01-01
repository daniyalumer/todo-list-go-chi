package service

import (
	"fmt"
	"time"

	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	repo "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateTodo(userID uint, body rq.Todo) (*dao.Todo, error) {
	var user dao.User

	if err := repo.FindByIdUser(&user, userID); err != nil {
		return &dao.Todo{}, fmt.Errorf("failed to find user: %v", err)
	}

	newTodo := dao.Todo{
		Description: body.Description,
		CompletedAt: nil,
		Completed:   false,
		UserID:      userID,
	}
	err := repo.Create(&newTodo)
	if err != nil {
		return &dao.Todo{}, fmt.Errorf("failed to create todo: %v", err)
	}
	return &newTodo, nil
}

func ReadTodoList() ([]dao.Todo, error) {
	var todos []dao.Todo

	err := repo.FindAll(&todos)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %v", err)
	}
	return todos, nil
}

func UpdateTodo(todoID uint, body rq.TodoUpdate) (dao.Todo, error) {
	var todo dao.Todo

	if err := repo.FindById(&todo, todoID); err != nil {
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

	if err := repo.Update(&todo, updates); err != nil {
		return dao.Todo{}, fmt.Errorf("failed to update todo: %v", err)
	}

	return todo, nil
}

func DeleteTodo(todoID uint) (dao.Todo, error) {
	var todo dao.Todo

	err := repo.FindById(&todo, todoID)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to find todo: %v", err)
	}

	err = repo.Delete(&todo)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to delete todo: %v", err)
	}

	return todo, nil
}
