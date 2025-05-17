package conectar

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// conectarnos a la BD
var Db *sql.DB

type ConfigDbEnv struct {
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_SERVER   string
	DB_PORT     string
}

func getDbEnv() ConfigDbEnv {
	return ConfigDbEnv{
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_SERVER:   os.Getenv("DB_SERVER"),
		DB_PORT:     os.Getenv("DB_PORT"),
	}
}

func Conectar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar las variables de entorno %v", err)
	}

	envDb := getDbEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", envDb.DB_USER, envDb.DB_PASSWORD, envDb.DB_SERVER, envDb.DB_PORT, envDb.DB_NAME)

	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	Db = connection
}

func Desconectar() {
	if Db != nil {
		Db.Close()
	}
}
