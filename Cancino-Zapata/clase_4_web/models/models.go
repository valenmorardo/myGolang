package models

type Cliente struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
}

type Clientes []Cliente

type Usuario struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
	Password string `json:"password"`
}
type Usuarios []Usuario

type UserLogin struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}
