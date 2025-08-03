package routes

import (
	//"fmt"
	"net/http"
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
