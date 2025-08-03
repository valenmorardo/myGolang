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
	"api_gin_bun/dto"

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

func PeliculasCreate(ctx *gin.Context) {
	var body dto.PeliculaDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}

	newPelicula := models.PeliculaModel{Nombre: body.Nombre, Descripcion: body.Descripcion, Year: body.Year, TematicaID: body.TematicaID}

	res, err := connection.DB.NewInsert().Model(&newPelicula).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}
	if rowsAffected, errRa := res.RowsAffected(); rowsAffected == 0 || errRa != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error. No se pudo insertar la pelicula",
			"error": errRa.Error(),
		})
		return
	}

	err = connection.DB.NewSelect().Model(&newPelicula).Relation("Tematica").Where("p.id = ?", newPelicula.ID).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Película creada, pero error al cargar temática",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Succesfull",
		"data": newPelicula,
	})
}

func PeliculasUpdate(ctx *gin.Context) {
	var body dto.PeliculaDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID Invalido",
			"info":  err.Error(),
		})
		return
	}

	// editar registro
	peliculaUpdated := models.PeliculaModel{
		ID: int64(id), Nombre: body.Nombre, Descripcion: body.Descripcion, Year: body.Year, TematicaID: body.TematicaID,
	}
	res, err := connection.DB.NewUpdate().Model(&peliculaUpdated).WherePK().Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}
	if rowsAffected, err := res.RowsAffected(); rowsAffected == 0 || err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error. No se pudo actualizar la pelicula",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Pelicula actualizada correctamente",
		"data": peliculaUpdated,
	})
}

func PeliculasDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error. ID Invalido",
			"error": err.Error(),
		})
		return
	}
	peliculaToDelete := models.PeliculaModel{ID: int64(id)}
	res, err := connection.DB.NewDelete().Model(&peliculaToDelete).WherePK().Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}
	ra, errRa := res.RowsAffected()
	if errRa != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error al verificar filas afectadas",
			"error": errRa.Error(),
		})
		return
	}
	if ra == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": fmt.Sprintf("No se encontró ninguna película con ID: %v", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("Pelicula con ID: %v, eliminada correctamente", id),
	})
}
