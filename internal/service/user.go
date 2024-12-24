package service

import (
	"fmt"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	"github.com/daniyalumer/todo-list-go-chi/internal/models"
	repository "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateUser(body rq.UserCreate) (models.User, error) {
	user := models.User{
		Username: body.Username,
	}
	err := repository.CreateUser(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func ReadUsers() ([]models.User, error) {
	var users []models.User
	err := repository.FindAllUsers(&users)
	if err != nil {
		return []models.User{}, fmt.Errorf("failed to read users: %v", err)
	}
	return users, nil
}

func DeleteUser(userID uint) (models.User, error) {
	tx := db.DB.Begin()

	var user models.User

	if err := repository.FindByIdUser(&user, userID); err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	if err := repository.DeleteUser(tx, &user); err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to delete user: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.User{}, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return user, nil
}
