package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HttpPort   string = "8080"
	DbHost     string = "localhost"
	DbPort     string = "6000"
	DbUser     string = "postgres"
	DbName     string = "postgres"
	DbPassword string
)

func Setup() error {
	err := godotenv.Load()

	if x := os.Getenv("HTTP_PORT"); x != "" {
		HttpPort = x
		log.Println("HTTP port configured to: ", HttpPort)
	} else {
		log.Println("Unable to configure HTTP port, defaulting to ", HttpPort)
	}

	if x := os.Getenv("DB_HOST"); x != "" {
		DbHost = x
		log.Println("DB Host configured to: ", DbHost)
	} else {
		log.Println("Unable to configure DB Host, defaulting to ", DbHost)
	}

	if x := os.Getenv("DB_PORT"); x != "" {
		DbPort = x
		log.Println("DB Port configured to: ", DbPort)
	} else {
		log.Println("Unable to configure DB Port, defaulting to ", DbPort)
	}

	if x := os.Getenv("DB_USER"); x != "" {
		DbUser = x
		log.Println("DB User configured to: ", DbUser)
	} else {
		log.Println("Unable to configure HTTP port, defaulting to ", DbUser)
	}

	if x := os.Getenv("DB_PASSWORD"); x != "" {
		DbPassword = x
		log.Println("DB Password configured")
	} else {
		log.Println("Unable to configure DB password")
	}

	if x := os.Getenv("DB_NAME"); x != "" {
		DbName = x
		log.Println("DB Name configured to: ", DbName)
	} else {
		log.Println("Unable to configure DB Name, defaulting to ", DbName)
	}

	return err
}
