package service

import (
	"fmt"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
)

func CreateUser(body rq.UserCreate) (models.User, error) {
	user := models.User{
		Username: body.Username,
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, fmt.Errorf("failed to create user: %v", result.Error)
	}
	return user, nil
}

func ReadUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Preload("Todos").Find(&users)
	if result.Error != nil {
		return []models.User{}, fmt.Errorf("failed to read users: %v", result.Error)
	}
	return users, nil
}

func DeleteUser(userID uint) (models.User, error) {
	tx := db.DB.Begin()

	var user models.User
	if err := tx.Preload("Todos").First(&user, userID).Error; err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	if err := tx.Where("user_id = ?", userID).Delete(&models.Todo{}).Error; err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to delete user's todos: %v", err)
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to delete user: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return user, nil
}
