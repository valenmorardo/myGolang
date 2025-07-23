// Package routes...
package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"api_gin_bun/dto"

	"github.com/gin-gonic/gin"
)

func EjemploGet(ctx *gin.Context) {
	ctx.Writer.Header().Set("Custom-Header", "I am a custom header!")
	auth := ctx.GetHeader("authorization")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metodo GET con GIN !!!!!!!",
		"Auth": auth,
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
		"data": body,
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
