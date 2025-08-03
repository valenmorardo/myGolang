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
	router.GET(routePrefix+"tematicas/:id", routes.TematicasGetByID)
	router.POST(routePrefix+"tematica-create", routes.TematicasCreate)
	router.PUT(routePrefix+"tematicas/:id", routes.TematicasEditByID)
	router.DELETE(routePrefix+"tematicas/:id", routes.TematicasDeleteByID)

	// router para peliculas (DB)
	router.GET(routePrefix+"peliculas", routes.PeliculasGet)
	router.GET(routePrefix+"peliculas/:id", routes.PeliculasGetByID)
	router.POST(routePrefix+"peliculas/create", routes.PeliculasCreate)
	router.PUT(routePrefix+"peliculas/:id",routes.PeliculasUpdate)
	router.DELETE(routePrefix+"peliculas/:id", routes.PeliculasDelete)

	fmt.Printf("\n\nCorriendo server de GIN en: localhost:%v\n\n", config.CfgEnv.SvPort)

	// le digo al server donde quiero que escuche, el puerto
	router.Run(":" + config.CfgEnv.SvPort) // listen and serve on 0.0.0.0:8080
}
