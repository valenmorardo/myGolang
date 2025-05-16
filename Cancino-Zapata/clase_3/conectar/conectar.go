package conectar

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

// conectarnos a la BD
var Db *sql.DB

func Conectar() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	print(errorVariables)
	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SEVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	Db = connection
}

func Desconectar() {
	Db.Close()
}
