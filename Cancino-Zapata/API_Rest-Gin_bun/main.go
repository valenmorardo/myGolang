package main

import (
	"fmt"

	"api_gin_bun/config"
	"api_gin_bun/connection"
	"api_gin_bun/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	connection.ConnectDB()
	defer connection.DB.Close()

	connection.Migrate() // ejecuto migrate para crear las tablas

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default() // doy inicio al router de gin
	routePrefix := "/api/"
	// static files
	router.Static("/public", "./public")
	// rutas de ejemplo
	router.GET(routePrefix+"get", routes.EjemploGet)
	router.POST(routePrefix+"post", routes.EjemploPost)
	router.PUT(routePrefix+"put/:id", routes.EjemploPut)
	router.DELETE(routePrefix+"delete", routes.EjemploDelete)
	// query & params routes example
	router.GET(routePrefix+"getParams/:id", routes.EjemploGetParams)
	router.GET(routePrefix+"queryString/", routes.EjemploGetQueryString)
	// upload file route example
	router.POST(routePrefix+"upload", routes.EjemploUpload)

	// router para tematicas (DB)
	router.GET(routePrefix+"tematicas", routes.TematicasGet)

	fmt.Printf("\n\nCorriendo server de GIN en: localhost:%v\n\n", config.CfgEnv.SvPort)

	// le digo al server donde quiero que escuche, el puerto
	router.Run(":" + config.CfgEnv.SvPort) // listen and serve on 0.0.0.0:8080
}
