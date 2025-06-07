package routes

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"clase_4_web/conectar"
	"clase_4_web/models"
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
