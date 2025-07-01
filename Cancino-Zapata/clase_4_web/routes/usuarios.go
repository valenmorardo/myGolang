package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"clase_4_web/conectar"
	"clase_4_web/middlewares"
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
	lastID, err := result.LastInsertId()
	if err != nil {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de registrar usuario.", nil, "ERROR")
		return
	}
	newUser.ID = int(lastID)
	utils.SendResponse(res, http.StatusOK, true, "Usuario creado correctamente", newUser, nil)
}

type userLogged struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}

func UserLogin(res http.ResponseWriter, req *http.Request) {
	userLoginData := models.UserLogin{}

	loginInput := req.Context().Value(middlewares.LoginInputKey)
	fmt.Println(loginInput)
	if loginInput == nil {
		utils.SendResponse(res, http.StatusBadRequest, false, "Error al obtener datos del usuario.", nil, "Error.")
		return
	}

	userLoginData, ok := loginInput.(models.UserLogin)
	if !ok {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Error interno al procesar login.", nil, "Error.")
		return
	}

	conectar.Conectar()
	defer conectar.Desconectar()

	sqlQuery := "SELECT id, nombre, correo, telefono, password FROM usuarios WHERE correo=?;"
	row := conectar.Db.QueryRow(sqlQuery, userLoginData.Correo)
	userFounded := models.Usuario{}
	err := row.Scan(&userFounded.ID, &userFounded.Nombre, &userFounded.Correo, &userFounded.Telefono, &userFounded.Password)
	if err == sql.ErrNoRows {
		utils.SendResponse(res, http.StatusUnauthorized, false, "Correo o contraseña inválidos", nil, nil)
		return
	} else if err != nil {
		utils.SendResponse(res, http.StatusInternalServerError, false, "Fallo a la hora de logear", nil, err.Error())
		return
	}
	fmt.Println(userLoginData)
	// comparar los hash pw
	passwordBytes := []byte(userLoginData.Password)
	passwordBytesDB := []byte(userFounded.Password)

	err = bcrypt.CompareHashAndPassword(passwordBytesDB, passwordBytes)
	if err != nil {
		utils.SendResponse(res, http.StatusUnauthorized, false, "Correo o contraseña inválidos", nil, nil)
		return
	}

	userLoggedRes := userLogged{
		ID:       userFounded.ID,
		Nombre:   userFounded.Nombre,
		Correo:   userFounded.Correo,
		Telefono: userFounded.Telefono,
	}
	fmt.Println("PASE EL todos los metodos de la ruta y ahora te doy la respuesta")
	utils.SendResponse(res, http.StatusOK, true, "Usuario Loggeado correctamente.", userLoggedRes, nil)
}
