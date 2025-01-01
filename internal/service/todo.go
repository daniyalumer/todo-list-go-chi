package service

import (
	"fmt"
	"time"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	repo "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateTodo(userID uint, body rq.Todo) (*dao.Todo, error) {
	var user dao.User

	DB := db.GetConnection()

	if err := repo.FindByIdUser(&user, DB, userID); err != nil {
		return &dao.Todo{}, fmt.Errorf("failed to find user: %v", err)
	}

	newTodo := dao.Todo{
		Description: body.Description,
		CompletedAt: nil,
		Completed:   false,
		UserID:      userID,
	}
	err := repo.Create(&newTodo, DB)
	if err != nil {
		return &dao.Todo{}, fmt.Errorf("failed to create todo: %v", err)
	}
	return &newTodo, nil
}

func ReadTodoList() ([]dao.Todo, error) {
	var todos []dao.Todo

	DB := db.GetConnection()

	err := repo.FindAll(&todos, DB)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %v", err)
	}
	return todos, nil
}

func UpdateTodo(todoID uint, body rq.TodoUpdate) (dao.Todo, error) {
	var todo dao.Todo

	DB := db.GetConnection()

	if err := repo.FindById(&todo, DB, todoID); err != nil {
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

	if err := repo.Update(&todo, DB, updates); err != nil {
		return dao.Todo{}, fmt.Errorf("failed to update todo: %v", err)
	}

	return todo, nil
}

func DeleteTodo(todoID uint) (dao.Todo, error) {
	var todo dao.Todo

	DB := db.GetConnection()

	err := repo.FindById(&todo, DB, todoID)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to find todo: %v", err)
	}

	err = repo.Delete(&todo, DB)
	if err != nil {
		return dao.Todo{}, fmt.Errorf("failed to delete todo: %v", err)
	}

	return todo, nil
}
