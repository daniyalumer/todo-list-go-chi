package db

import (
	"fmt"
	"log"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbName,
		conf.DbPassword,
	)

	var err error
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	return err
}

func GetConnection() *gorm.DB {
	return conn
}
