package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HttpPort string = "8080"
)

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if x := os.Getenv("HTTP_PORT"); x != "" {
		HttpPort = x
		log.Println("HTTP port configured to: ", HttpPort)
	} else {
		log.Println("Unable to configure HTTP port, defaulting to ", HttpPort)
	}
}
