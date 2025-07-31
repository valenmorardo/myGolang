// Package models...
package models

import (
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

type TematicaModel struct {
	bun.BaseModel `bun:"table:tematicas,alias:t"`

	ID     int64  `bun:",pk,autoincrement" json:"id"`
	Nombre string `bun:"nombre,notnull" json:"nombre"`
	Slug   string `bun:"slug,notnull" json:"slug"`
}

type PeliculaModel struct {
	bun.BaseModel `bun:"table:peliculas,alias:p"`

	ID          int64  `bun:",pk,autoincrement" json:"id"`
	Nombre      string `bun:"nombre,notnull" json:"nombre"`
	Descripcion string `bun:"descripcion" json:"descripcion"`
	Year        int    `bun:"year" json:"year"`

	TematicaID int64         `bun:"tematica_id" json:"tematicaID"`
	Tematica   TematicaModel `bun:"rel:belongs-to,join:tematica_id=id" json:"tematica"`
}
