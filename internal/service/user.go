package service

import (
	"fmt"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	repository "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateUser(req rq.UserCreate) (dao.User, error) {
	user := dao.User{
		Username: req.Username,
	}
	err := repository.CreateUser(&user)
	if err != nil {
		return dao.User{}, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func ReadUsers() ([]dao.User, error) {
	var users []dao.User
	err := repository.FindAllUsers(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to read users: %v", err)
	}
	return users, nil
}

func DeleteUser(userID uint) (dao.User, error) {
	tx := db.DB.Begin()

	var user dao.User

	if err := repository.FindByIdUser(&user, userID); err != nil {
		tx.Rollback()
		return dao.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	if err := repository.DeleteUser(tx, &user); err != nil {
		tx.Rollback()
		return dao.User{}, fmt.Errorf("failed to delete user: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return dao.User{}, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return user, nil
}
