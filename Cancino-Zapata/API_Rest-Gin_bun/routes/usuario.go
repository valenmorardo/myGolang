package routes

import (
	"net/http"

	"api_gin_bun/connection"
	"api_gin_bun/dto"
	"api_gin_bun/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UsuarioRegister(ctx *gin.Context) {
	var body dto.UsuarioDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}

	exists, err := connection.DB.NewSelect().Model((*models.UserModel)(nil)).Where("email=?", body.Email).Exists(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "El email ya existe",
		})
		return
	}

	// hasheamos el pw
	costo := 10
	bytesHashedPw, err := bcrypt.GenerateFromPassword([]byte(body.Password), costo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}
	newUserModel := models.UserModel{Nombre: body.Nombre, Email: body.Email, Password: string(bytesHashedPw)}

	res, err := connection.DB.NewInsert().Model(&newUserModel).Exec(ctx.Request.Context())
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "No se pudo registrar al usuario",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":  "Usuario registrado con exito",
		"user": newUserModel,
	})
}
