package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"api_gin_bun/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	routePrefix := "/api/"

	router := gin.Default() // doy inicio al router de gin

	// ruta get de ejemplo
	router.GET(routePrefix+"get", routes.EjemploGet)
	router.POST(routePrefix+"post", routes.EjemploPost)
	router.PUT(routePrefix+"put/:id",routes.EjemploPut)
	router.DELETE(routePrefix+"delete", routes.EjemploDelete)

	router.GET(routePrefix+"getParams/:id", routes.EjemploGetParams)
	router.GET(routePrefix+"queryString/", routes.EjemploGetQueryString)

	router.POST(routePrefix+"upload", routes.EjemploUpload)

	// variables de entorno
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("\n\nCorriendo server de GIN en: localhost:%v\n\n", os.Getenv("SV_PORT") )

	// le digo al server donde quiero que escuche, el puerto
	router.Run(":" + os.Getenv("SV_PORT")) // listen and serve on 0.0.0.0:8080
	
}
