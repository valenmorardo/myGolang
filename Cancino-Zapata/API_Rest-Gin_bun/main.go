package main

import (
	"fmt"

	"api_gin_bun/config"
	"api_gin_bun/connection"
	"api_gin_bun/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := connection.ConnectDB()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	routePrefix := "/api/"

	router := gin.Default() // doy inicio al router de gin

	// static files
	router.Static("/public", "./public")
	// ruta get de ejemplo
	router.GET(routePrefix+"get", routes.EjemploGet)
	router.POST(routePrefix+"post", routes.EjemploPost)
	router.PUT(routePrefix+"put/:id", routes.EjemploPut)
	router.DELETE(routePrefix+"delete", routes.EjemploDelete)

	router.GET(routePrefix+"getParams/:id", routes.EjemploGetParams)
	router.GET(routePrefix+"queryString/", routes.EjemploGetQueryString)

	router.POST(routePrefix+"upload", routes.EjemploUpload)

	fmt.Printf("\n\nCorriendo server de GIN en: localhost:%v\n\n", config.CfgEnv.SvPort)

	// le digo al server donde quiero que escuche, el puerto
	router.Run(":" + config.CfgEnv.SvPort) // listen and serve on 0.0.0.0:8080
}
