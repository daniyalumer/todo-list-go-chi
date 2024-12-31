package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Db *sql.DB

func Connect() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbName,
		conf.DbPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
	Db, err = db.DB()
	if err != nil {
		log.Printf("failed to connect to the database: %v", err)
	}
	log.Println("Database connected successfully")

	return err
}
