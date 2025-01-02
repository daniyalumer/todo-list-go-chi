package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Todos    []Todo `json:"todos" gorm:"foreignKey:UserID"`
}
