// Package models...
package models

import (
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

type TematicasModel struct {
	bun.BaseModel `bun:"table:tematicas,alias:t"`

	ID     int64  `bun:",pk,autoincrement" json:"id"`
	Nombre string `bun:"nombre,notnull" json:"nombre"`
	Slug   string `bun:"slug,notnull" json:"slug"`
}
