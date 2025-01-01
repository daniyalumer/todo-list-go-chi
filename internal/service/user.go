package service

import (
	"fmt"

	"github.com/daniyalumer/todo-list-go-chi/db/dao"
	"github.com/daniyalumer/todo-list-go-chi/internal/http/rq"
	repo "github.com/daniyalumer/todo-list-go-chi/internal/repo"
)

func CreateUser(req rq.User) (dao.User, error) {
	user := dao.User{
		Username: req.Username,
	}
	err := repo.CreateUser(&user)
	if err != nil {
		return dao.User{}, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func ReadUsers() ([]dao.User, error) {
	var users []dao.User

	err := repo.FindAllUsers(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to read users: %v", err)
	}
	return users, nil
}

func DeleteUser(userID uint) (dao.User, error) {
	var user dao.User

	if err := repo.FindByIdUser(&user, userID); err != nil {
		return dao.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	if err := repo.DeleteUser(&user); err != nil {
		return dao.User{}, fmt.Errorf("failed to delete user: %v", err)
	}

	return user, nil
}
