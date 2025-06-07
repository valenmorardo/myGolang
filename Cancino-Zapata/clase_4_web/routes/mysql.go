package routes

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"net/http"

	"clase_4_web/conectar"
	"clase_4_web/models"
	"clase_4_web/utils"
	//"clase_4_web/utils"
)

func Mysql_get(res http.ResponseWriter, req *http.Request) {
	// conexion a bd
	conectar.Conectar()
	defer conectar.Desconectar()

	sqlQuery := "SELECT id, nombre, correo, telefono FROM clientes order by id desc"
	clientes := models.Clientes{}

	datos, err := conectar.Db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer datos.Close()

	for datos.Next() {
		cliente := models.Cliente{}
		err := datos.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
		if err != nil {
			http.Error(res, "Error al leer datos.", http.StatusInternalServerError)
		}
		clientes = append(clientes, cliente)
	}
	response := map[string]any{
		"success": true,
		"data":    clientes,
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(response)
}

func Mysql_create(res http.ResponseWriter, req *http.Request) {
	conectar.Conectar()
	defer conectar.Desconectar()

	newClient := models.Cliente{}

	err := json.NewDecoder(req.Body).Decode(&newClient)
	if err != nil {
		http.Error(res, "error al recibir datos.", http.StatusInternalServerError)
		return
	}

	sqlQuery := "INSERT into clientes (nombre, correo, telefono) values (?, ?, ?);"
	result, err := conectar.Db.Exec(sqlQuery, newClient.Nombre, newClient.Correo, newClient.Telefono)
	if err != nil {
		fmt.Println("error al crear el cliente: ", err)
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error al obtener el id del cliente.", err)
	}
	newClient.Id = int(lastId)

	utils.SendResponse(res, http.StatusOK, true, "Cliente registrado correctamente", newClient, nil)
}
