package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {

	// err := godotenv.Load()
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, using system environment")
	// }
	// if err != nil {
	// 	log.Fatal("error loading .env file:", err.Error())
	// }
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	expirent, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Tz:       os.Getenv("DB_TZ"),
		},
		Jwt: Jwt{
			Key:    os.Getenv("JWT_KEY"),
			Expire: expirent,
		},
	}
}
