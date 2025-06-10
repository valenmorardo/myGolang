package routes

import (
	"encoding/json"
	"net/http"

	"clase_4_web/conectar"
	"clase_4_web/models"
	"clase_4_web/utils"

	"golang.org/x/crypto/bcrypt"
)

func isUserDataValid(user models.Usuario) bool {
	if user.Correo == "" || user.Nombre == "" || user.Telefono == "" || user.Password == "" {
		return false
	}
	return true
}

func UserRegister(res http.ResponseWriter, req *http.Request) {
	conectar.Conectar()
	defer conectar.Desconectar()

	newUser := models.Usuario{}
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		utils.SendResponse(res, http.StatusBadRequest, false, "Error al obtener datos del usuario.", nil, err.Error())
		return
	}
	if !isUserDataValid(newUser) {
		utils.SendResponse(res, http.StatusBadRequest, false, "ERROR: Todos los campos son obligatorios.", nil, "Error.")
		return
	}

	// bcrypt
	costo := 8
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), costo)
	if err != nil {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de registrar usuario.", nil, err.Error())
		return
	}

	sqlQuery := "INSERT INTO usuarios (nombre, correo, telefono, password) VALUES(?,?,?,?);"

	result, err := conectar.Db.Exec(sqlQuery, newUser.Nombre, newUser.Correo, newUser.Telefono, string(hashedPassword))
	if err != nil {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de registrar usuario.", nil, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de registrar usuario.", nil, err.Error())
		return
	}

	if rowsAffected == 0 {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de registrar usuario.", nil, "ERROR")
		return
	}
	newUser.Password = string(hashedPassword)
	lastId, err := result.LastInsertId()
	if err != nil {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de registrar usuario.", nil, "ERROR")
		return
	}
	newUser.Id = int(lastId)
	utils.SendResponse(res, http.StatusOK, true, "Usuario creado correctamente", newUser, nil)
}
