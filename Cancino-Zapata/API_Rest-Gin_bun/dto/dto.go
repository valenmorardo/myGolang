// Package dto...
package dto

type UserDto struct {
	Nombre string `json:"nombre"`
}

type TematicaDto struct {
	Nombre string `json:"nombre" binding:"required"`
}

type PeliculaDto struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	TematicaID  int64  `json:"tematicaID" binding:"required"`
}
