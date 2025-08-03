package routes

import (
	"net/http"

	"api_gin_bun/connection"
	"api_gin_bun/dto"
	"api_gin_bun/jwt"
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

func UsuarioLogin(ctx *gin.Context) {
	var bodyLogin dto.UsuarioLoginDto
	if err := ctx.ShouldBindJSON(&bodyLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Error.",
			"error": err.Error(),
		})
		return
	}

	usuarioModel := models.UserModel{}
	err := connection.DB.NewSelect().Model(&usuarioModel).Where("email=?", bodyLogin.Email).Limit(1).Scan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "las credenciales son incorrectas",
			"error": err.Error(),
		})
		return
	}

	errCheckPw := bcrypt.CompareHashAndPassword([]byte(usuarioModel.Password), []byte(bodyLogin.Password))
	if errCheckPw != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "las credenciales ingresadas son incorrectas.",
		})
		return
	}

	jwtKey, jwtErr := jwt.JWTGenerator(usuarioModel.Email, usuarioModel.Nombre, usuarioModel.ID)
	if jwtErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "ocurrio un error al generar jwt",
			"error": jwtErr.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "usuario loggeado correctamente.",
		"user": gin.H{
			"id":     usuarioModel.ID,
			"nombre": usuarioModel.Nombre,
			"email":  usuarioModel.Email,
		},
		"jwtToken": jwtKey,
	})
}
