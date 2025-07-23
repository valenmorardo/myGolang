// Package routes...
package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"api_gin_bun/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EjemploGet(ctx *gin.Context) {
	ctx.Writer.Header().Set("Custom-Header", "I am a custom header!")
	auth := ctx.GetHeader("authorization")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metodo GET con GIN !!!!!!!",
		"Auth":    auth,
	})
}

func EjemploPost(ctx *gin.Context) {
	var body dto.UserDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error. Bad Request",
			"error":   err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metodo post con GIN",
		"data":    body,
	})
}

func EjemploPut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metodo put con GIN",
	})
}

func EjemploDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "metodo delete con GIN!",
	})
}

func EjemploGetParams(ctx *gin.Context) {
	fmt.Println()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   "ID Invalido",
			"message": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metodo GET with PARAMS",
		"param":   id,
	})
}

func EjemploGetQueryString(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":   "ID Invalido",
			"message": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metodo GET with PARAMS",
		"param":   id,
	})
}

func EjemploUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("foto")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "archivo no válido", "detail": err.Error(),
		})
		return
	}

	ext := filepath.Ext(file.Filename) // incluye el punto: ".jpg"
	filename := uuid.New().String() + ext
	dst := filepath.Join("public", "uploads", "fotos", filename)

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo guardar el archivo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "foto guardada correctamente",
		"filename": filename,
	})
}
