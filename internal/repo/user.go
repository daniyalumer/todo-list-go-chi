package repo

import (
	"github.com/daniyalumer/todo-list-go-chi/db"
	"github.com/daniyalumer/todo-list-go-chi/db/dao"
)

func CreateUser(user *dao.User) error {
	DB := db.GetConnection()
	return DB.Create(&user).Error
}

func FindAllUsers(users *[]dao.User) error {
	DB := db.GetConnection()
	return DB.Preload("Todos").Find(&users).Error
}

func FindByIdUser(user *dao.User, userID uint) error {
	DB := db.GetConnection()
	return DB.Preload("Todos").Find(&user, userID).Error
}

func DeleteUser(user *dao.User) error {
	DB := db.GetConnection()
	if err := DB.Where("user_id = ?", user.ID).Delete(&dao.Todo{}).Error; err != nil {
		return err
	}

	if err := DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckDeleted(user *dao.User, userID uint) error {
	DB := db.GetConnection()
	return DB.Where("deleted_at IS NULL").Preload("Todos").First(user, userID).Error
}
