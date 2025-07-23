package routes

import (
	"context"

	"net/http"

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
