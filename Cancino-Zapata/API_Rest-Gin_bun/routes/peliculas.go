package routes

import (
	//"context"
	//"database/sql"
	//"errors"

	//	"fmt"

	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"api_gin_bun/connection"
	//"api_gin_bun/dto"
	"api_gin_bun/models"

	"github.com/gin-gonic/gin"
)

func PeliculasGet(ctx *gin.Context) {
	var peliculas []models.PeliculaModel

	err := connection.DB.NewSelect().Model(&peliculas).Relation("Tematica").Order("id DESC").Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al obtener peliculas",
			"info":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, peliculas)
}

func PeliculasGetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error. ID Invalido",
			"error": err.Error(),
		})
		return
	}

	var pelicula models.PeliculaModel
	err = connection.DB.NewSelect().Model(&pelicula).Where("p.id=?", id).Relation("Tematica").Scan(ctx.Request.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": fmt.Sprintf("No se encontró la película con ID %d", id),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error al consultar DB.",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Pelicula encontrada",
		"data": pelicula,
	})
}
