package routes

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"api_gin_bun/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"api_gin_bun/connection"
)

func PeliculaFotoUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("foto")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "archivo no válido", "detail": err.Error(),
		})
		return
	}

	ext := filepath.Ext(file.Filename) // incluye el punto: ".jpg"
	filename := uuid.New().String() + ext
	dst := filepath.Join("public", "uploads", "peliculas", filename)

	if err = ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo guardar el archivo"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error. ID Invalido",
			"error": err.Error(),
		})
		return
	}

	fotoModel := models.PeliculaFotoModel{Nombre: filename, PeliculaID: int64(id)}
	res, err := connection.DB.NewInsert().Model(&fotoModel).Exec(ctx.Request.Context())
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
			"msg": "No se logro almacenar la foto de la pelicula.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "foto de la pelicula guardada correctamente",
		"filename": filename,
	})
}

func PeliculaFotoGetByID(ctx *gin.Context) {
	idPelicula, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID Invalido",
			"info":  err.Error(),
		})
		return
	}

	var fotosPelicula []models.PeliculaFotoModel

	err = connection.DB.NewSelect().Model(&fotosPelicula).Where("pelicula_id=?", idPelicula).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}
	if len(fotosPelicula) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": fmt.Sprintf("No se encontraron fotos para la película con ID %v", idPelicula),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"datos": fotosPelicula,
		"msg":   "Fotos encontradas",
	})
}

func PeliculaFotoDeleteByID(ctx *gin.Context) {
	idPelicula, err := strconv.Atoi(ctx.Param("idP"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID no válido", "detail": err.Error(),
		})
		return
	}
	idFoto, err := strconv.Atoi(ctx.Param("idF"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID no válido", "detail": err.Error(),
		})
		return
	}

	fotoPeliculaToDelete := models.PeliculaFotoModel{ID: int64(idFoto), PeliculaID: int64(idPelicula)}
	err = connection.DB.NewSelect().Model(&fotoPeliculaToDelete).WherePK().Where("pelicula_id=?", idPelicula).Limit(1).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg":   "No se encontró la foto",
			"error": err.Error(),
		})
		return
	}

	res, err := connection.DB.NewDelete().Model(&fotoPeliculaToDelete).WherePK().Exec(ctx.Request.Context())
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
			"msg": "No se logro eliminar la foto de la pelicula.",
		})
		return
	}

	rutaArchivo := filepath.Join("public", "uploads", "peliculas", fotoPeliculaToDelete.Nombre)
	if err := os.Remove(rutaArchivo); err != nil && !errors.Is(err, os.ErrNotExist) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error al eliminar el archivo del disco",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "foto de la pelicula eliminada correctamente",
	})
}
