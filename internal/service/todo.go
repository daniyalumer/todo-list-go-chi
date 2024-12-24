package service

import (
	"fmt"
	"time"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func CreateTodo(userid uint, body rq.TodoCreate) (models.Todo, error) {
	var user models.User

	if err := db.DB.First(&user, userid).Error; err != nil {
		return models.Todo{}, fmt.Errorf("failed to find user: %v", err)
	}

	newTodo := models.Todo{
		Description: body.Description,
		CompletedAt: nil,
		Completed:   false,
		UserID:      userid,
	}
	result := db.DB.Create(&newTodo)
	if result.Error != nil {
		return models.Todo{}, fmt.Errorf("failed to create todo: %v", result.Error)
	}
	return newTodo, nil
}

func ReadTodoList() ([]models.Todo, error) {
	var todos []models.Todo
	results := db.DB.Find(&todos)
	if results.Error != nil {
		return nil, fmt.Errorf("failed to fetch todos: %v", results.Error)
	}
	return todos, nil
}

func UpdateTodo(todoID uint, body rq.TodoUpdate) (models.Todo, error) {
	var todo models.Todo
	if result := db.DB.First(&todo, todoID); result.Error != nil {
		return models.Todo{}, fmt.Errorf("todo not found: %v", result.Error)
	}
	updates := map[string]interface{}{}
	if body.Description != "" {
		updates["description"] = body.Description
	}

	if body.Completed {
		updates["completed"] = body.Completed
		updates["completed_at"] = time.Now()
	}

	if err := db.DB.Model(&todo).Updates(updates).Error; err != nil {
		return models.Todo{}, fmt.Errorf("failed to update todo: %v", err)
	}

	return todo, nil
}

func DeleteTodo(todoID uint) (models.Todo, error) {
	var todo models.Todo

	results := db.DB.First(&todo, todoID)
	if results.Error != nil {
		return models.Todo{}, fmt.Errorf("failed to find todo: %v", results.Error)
	}

	results = db.DB.Delete(&todo)
	if results.Error != nil {
		return models.Todo{}, fmt.Errorf("failed to delete todo: %v", results.Error)
	}
	return todo, nil
}
