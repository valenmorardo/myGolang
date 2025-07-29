// Package dto...
package dto

type UserDto struct {
	Nombre string `json:"nombre"`
}

type TematicaDto struct {
	Nombre string `json:"nombre" binding:"required"`
}
