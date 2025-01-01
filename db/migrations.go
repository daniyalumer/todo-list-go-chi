package db

import (
	"database/sql"
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var sqlFiles embed.FS

const MigrationVersion = 1

func RunMigrations() {
	Db := GetConnection()

	DB, err := Db.DB()
	if err != nil {
		log.Panicf("failed to access underlying database connection: %v", err)
	}
	// defer DB.Close()

	log.Println("Database connection established")

	m := createMigrateInstance(DB)
	log.Printf("Looking for migrations in: file://db/migrations/")

	if err := m.Migrate(MigrationVersion); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply")
			return
		}
		log.Fatalf("failed to run migration: %v", err)
	}

	log.Println("Migration applied successfully!")
}

func DownMigrations() {
	Db := GetConnection()

	db, err := Db.DB()
	if err != nil {
		log.Panicf("failed to access underlying database connection: %v", err)
	}
	// defer DB.Close()

	log.Println("Database connection established")

	m := createMigrateInstance(db)
	log.Printf("Looking for migrations in: file://db/migrations/")

	if err := m.Down(); err != nil {
		log.Fatalf("failed to run migration: %v", err)
	}

	log.Println("Migration rolled back successfully!")
}

func createMigrateInstance(db *sql.DB) *migrate.Migrate {
	dirInstance, err := iofs.New(sqlFiles, "migrations")
	if err != nil {
		log.Panicf("failed to create directory instance: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panicf("failed to create driver: %v", err)
	}

	m, err := migrate.NewWithInstance("iofs", dirInstance, "postgres", driver)
	if err != nil {
		log.Panicf("failed to create migrate instance: %v", err)
	}
	return m
}
