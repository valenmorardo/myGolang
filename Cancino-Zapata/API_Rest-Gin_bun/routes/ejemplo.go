// Package routes...
package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func EjemploGet(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Metodo GET con GIN !!!",
	})
}

func EjemploPost(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Metodo post con GIN",
	})
}

func EjemploPut(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Metodo put con GIN",
	})
}

func EjemploDelete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "metodo delete con GIN!",
	})
}

func EjemploGetParams(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error":   "ID Invalido",
			"message": err,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Metodo GET with PARAMS",
		"param":   id,
	})
}
