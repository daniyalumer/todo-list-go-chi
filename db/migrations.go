package db

import (
	"database/sql"
	"embed"
	"log"

	"github.com/daniyalumer/todo-list-go-chi/conf"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var sqlFiles embed.FS

const MigrationVersion = 1

func RunMigrations() {
	db := connectDatabase()
	defer db.Close()

	log.Println("Database connection established")

	m := createMigrateInstance(db)
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

func RollbackMigrations() {
	db := connectDatabase()
	defer db.Close()

	m := createMigrateInstance(db)

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to rollback migration: %v", err)
	}

	log.Println("Migration rolled back successfully!")
}

func connectDatabase() *sql.DB {
	db, err := sql.Open("postgres", "host="+conf.DbHost+" port="+conf.DbPort+" user="+conf.DbUser+" dbname="+conf.DbName+" password="+conf.DbPassword+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return db
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
