// Package connection....
package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"api_gin_bun/config"
)	

var DB *bun.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.CfgEnv.DBUser,
		config.CfgEnv.DBPassword,
		config.CfgEnv.DBName,
		config.CfgEnv.DBHost,
		config.CfgEnv.DBPort,
	)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir conexión: %v", err)
	}

	// Probamos si realmente conecta
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error al conectar a la DB: %v", err)
	}

	DB = bun.NewDB(sqlDB, pgdialect.New())
}