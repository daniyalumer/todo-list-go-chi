package service

import (
	"fmt"

	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	repository "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateUser(req rq.User) (dao.User, error) {
	DB, err := db.Connect()
	if err != nil {
		return dao.User{}, fmt.Errorf("failed to connect to database: %v", err)
	}

	user := dao.User{
		Username: req.Username,
	}
	err = repository.CreateUser(&user, DB)
	if err != nil {
		return dao.User{}, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func ReadUsers() ([]dao.User, error) {
	var users []dao.User

	DB, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = repository.FindAllUsers(&users, DB)
	if err != nil {
		return nil, fmt.Errorf("failed to read users: %v", err)
	}
	return users, nil
}

func DeleteUser(userID uint) (dao.User, error) {
	DB, err := db.Connect()
	if err != nil {
		return dao.User{}, fmt.Errorf("failed to connect to database: %v", err)
	}

	var user dao.User

	if err := repository.FindByIdUser(&user, DB, userID); err != nil {
		return dao.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	if err := repository.DeleteUser(DB, &user); err != nil {
		return dao.User{}, fmt.Errorf("failed to delete user: %v", err)
	}

	return user, nil
}
