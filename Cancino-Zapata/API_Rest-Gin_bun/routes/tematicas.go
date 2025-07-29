package routes

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"api_gin_bun/connection"
	"api_gin_bun/models"

	"github.com/gin-gonic/gin"
)

func TematicasGet(ctx *gin.Context) {
	var tematicas []models.TematicaModel

	err := connection.DB.NewSelect().Model(&tematicas).Order("id DESC").Scan(context.TODO())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al obtener tematicas",
			"info":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, tematicas)
}

func TematicasGetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID Invalido",
			"info":  err.Error(),
		})
		return
	}

	var tematica models.TematicaModel
	err = connection.DB.NewSelect().Model(&tematica).Where("id=?", id).Scan(context.TODO())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "No se encontro la tematica",
				"infoErr": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error al consultar la base de datos",
				"info":  err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Temática encontrada",
		"data": tematica,
	})
}
