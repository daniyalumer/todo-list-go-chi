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
	Db, err := Connect()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	DB, err := Db.DB()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer DB.Close()

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

func createMigrateInstance(db *sql.DB) *migrate.Migrate {
	dirInstance, err := iofs.New(sqlFiles, "migrations")
	if err != nil {
		log.Fatalf("failed to create directory instance: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create driver: %v", err)
	}

	m, err := migrate.NewWithInstance("iofs", dirInstance, "postgres", driver)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	return m
}
