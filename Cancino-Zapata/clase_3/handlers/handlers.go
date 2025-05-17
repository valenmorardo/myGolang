package handlers

import (
	"database/sql"
	"fmt"

	"clase_3_mysql_driver/conectar"
	"clase_3_mysql_driver/models"
)

func GetAll() {
	conectar.Conectar()

	sqlQuery := "select id, nombre, correo, telefono from clientes order by id desc;"

	datos, err := conectar.Db.Query(sqlQuery)
	if err != nil {
		fmt.Println(err)
	}

	defer conectar.Desconectar()
	defer datos.Close()

	for datos.Next() {
		dato := models.Cliente{}

		datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		fmt.Printf("Id: %v | Nombre: %v | Correo: %v | Telefono: %v\n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)

	}
}

func GetById(id int) {
	conectar.Conectar()
	defer conectar.Desconectar()

	sqlQuery := "select id, nombre, correo, telefono from clientes where id=?;"
	fila := conectar.Db.QueryRow(sqlQuery, id)

	cliente := models.Cliente{}
	err := fila.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no se encontraron reigstros con este ID")
		}
		fmt.Println("Eror al scanear fila de bd: ", err)
		return
	}
	fmt.Println(cliente)
}

func CreateClient(cliente models.Cliente) {

	conectar.Conectar()
	defer conectar.Desconectar()

	sql := "insert into clientes (nombre, correo, telefono) values (?, ?, ?);"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		fmt.Println("error al crear el cliente: ", err)
	}

	fmt.Println("Nuevo cliente registrado con exito. ", result)
}
