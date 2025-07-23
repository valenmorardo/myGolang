// Package connection....
package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL

	"api_gin_bun/config"
)	

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.CfgEnv.DBUser,
		config.CfgEnv.DBPassword,
		config.CfgEnv.DBName,
		config.CfgEnv.DBHost,
		config.CfgEnv.DBPort,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir conexión: %v", err)
	}

	// Probamos si realmente conecta
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al conectar a la DB: %v", err)
	}

	return db
}