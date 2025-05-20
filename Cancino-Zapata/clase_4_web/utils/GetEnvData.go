package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvData struct {
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_SERVER   string
	DB_PORT     string
	SERVER_PORT string
}

func GetEnvData() EnvData {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error al cargar archivo .env:")
		panic(err)
	}
	return EnvData{
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_SERVER:   os.Getenv("DB_SERVER"),
		DB_PORT:     os.Getenv("DB_PORT"),
		SERVER_PORT: os.Getenv("SERVER_PORT"),
	}
}
