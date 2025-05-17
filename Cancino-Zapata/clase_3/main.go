package main

import (
	"clase_3_mysql_driver/handlers"
	"clase_3_mysql_driver/models"
)

func main(){
	/* handlers.GetById(12) */

	niufa := models.Cliente {
		Nombre: "Niufa",
		Correo: "niufa@gmail.com",
		Telefono: "21321321",
	}
	handlers.CreateClient(niufa)

}