package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	//"fmt"
	"net/http"

	"clase_4_web/conectar"
	"clase_4_web/models"
	"clase_4_web/utils"

	"github.com/gorilla/mux"
	//"clase_4_web/utils"
)

func Mysql_get(res http.ResponseWriter, req *http.Request) {
	// conexion a bd
	conectar.Conectar()
	defer conectar.Desconectar()

	sqlQuery := "SELECT ID, nombre, correo, telefono FROM clientes order by id desc"
	clientes := models.Clientes{}

	datos, err := conectar.Db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer datos.Close()

	for datos.Next() {
		cliente := models.Cliente{}
		err := datos.Scan(&cliente.ID, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
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

	lastID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error al obtener el id del cliente.", err)
	}
	newClient.ID = int(lastID)

	utils.SendResponse(res, http.StatusOK, true, "Cliente registrado correctamente", newClient, nil)
}

func Mysql_editar(res http.ResponseWriter, req *http.Request) {
	conectar.Conectar()
	defer conectar.Desconectar()

	clientIdToEdit := mux.Vars(req)["id"]

	newDataClient := models.Cliente{}

	err := json.NewDecoder(req.Body).Decode(&newDataClient)
	if err != nil {
		http.Error(res, "Error al obtener datos del cliente.", http.StatusInternalServerError)
	}

	sqlQuery := "UPDATE clientes set nombre=?, correo=?, telefono=? where id=?;"
	_, err = conectar.Db.Exec(sqlQuery, newDataClient.Nombre, newDataClient.Correo, newDataClient.Telefono, clientIdToEdit)
	if err != nil {
		fmt.Println("Error al crear cliente.", err)
		return
	}
	newDataClient.ID, err = strconv.Atoi(clientIdToEdit)
	if err != nil {
		http.Error(res, "ID Invalido.", http.StatusBadRequest)
	}
	utils.SendResponse(res, http.StatusOK, true, "Cliente actualizado correctamente.", newDataClient, nil)
}

func Mysql_delete(res http.ResponseWriter, req *http.Request) {
	conectar.Conectar()
	defer conectar.Desconectar()

	clientIDToDelete := mux.Vars(req)["id"]

	sqlQuery := "DELETE FROM clientes where id=?;"
	result, err := conectar.Db.Exec(sqlQuery, clientIDToDelete)
	if err != nil {
		http.Error(res, "Error al eliminar cliente", http.StatusInternalServerError)
	}
	rowsAffec, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffec == 0 {
		utils.SendResponse(res, http.StatusBadRequest, true, "No se encontro un cliente con ese ID", nil, nil)
		return
	}
	msgRes := fmt.Sprintf("Se elimino el cliente con id: %v", clientIDToDelete)
	utils.SendResponse(res, http.StatusOK, true, msgRes, nil, nil)
}
