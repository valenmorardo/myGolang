package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	routePrefix := "/api/"

	router := gin.Default() // doy inicio al router de gin

	// ruta get de ejemplo
	router.GET(routePrefix+"hola", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hola desde GIN !!!",
		})
		
	})
	// variables de entorno
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("\n\nCorriendo server de GIN en: localhost:%v\n\n", os.Getenv("SV_PORT") )

	// le digo al server donde quiero que escuche, el puerto
	router.Run(":" + os.Getenv("SV_PORT")) // listen and serve on 0.0.0.0:8080
	
}
