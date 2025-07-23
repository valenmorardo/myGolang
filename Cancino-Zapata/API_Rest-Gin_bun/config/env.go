// Package config....
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvStruct struct {
	SvPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

var CfgEnv *EnvStruct

func init() {
	// carga el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	// inicializa la configuración
	CfgEnv = &EnvStruct{ // creo una instancia de EnvStruct y le asigno el puntero a cfgenv para usarlo en todo el proyecto
		SvPort:     os.Getenv("SV_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
