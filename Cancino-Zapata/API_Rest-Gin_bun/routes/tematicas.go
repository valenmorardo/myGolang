package routes

import (
	"context"
	"database/sql"
	"errors"

	//	"fmt"
	"net/http"
	"strconv"

	"api_gin_bun/connection"
	"api_gin_bun/dto"
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
				"error":   "No se encontro la tematica",
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

func TematicasCreate(ctx *gin.Context) {
	var body dto.TematicaDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}

	newTematica := models.TematicaModel{Nombre: body.Nombre, Slug: body.Nombre + "-slug"}

	_, err := connection.DB.NewInsert().Model(&newTematica).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Succesfull",
		"data": newTematica,
	})
}

func TematicasEditByID(ctx *gin.Context) {
	var body dto.TematicaDto

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

	// edtiar registro
	tematicaModifed := models.TematicaModel{ID: int64(id), Nombre: body.Nombre, Slug: body.Nombre + "-"}
	res, err := connection.DB.NewUpdate().Model(&tematicaModifed).WherePK().Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "ERROR. No se encontro tematica con ese ID.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Successfull. Tematica modificada",
		"data": tematicaModifed,
	})
}

func TematicasDeleteByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID Invalido",
			"info":  err.Error(),
		})
		return
	}
	res, err := connection.DB.NewDelete().Model((*models.TematicaModel)(nil)).Where("id=?", id).Exec(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error",
			"error": err.Error(),
		})
		return
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "ERROR. No se encontro tematica con ese ID.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Successfull. Tematica eliminada",
	})
}
