package connection

import (
	"context"

	"api_gin_bun/models"
)

func Migrate() {
	DB.NewCreateTable().
		Model((*models.TematicaModel)(nil)).
		IfNotExists().
		Exec(context.TODO())

	DB.NewCreateTable().
		Model((*models.PeliculaModel)(nil)).
		ForeignKey(`("tematica_id") REFERENCES "tematicas"("id") ON DELETE CASCADE`).
		IfNotExists().
		Exec(context.TODO())

	DB.NewCreateTable().
		Model((*models.PeliculaFotoModel)(nil)).
		ForeignKey(`("pelicula_id") REFERENCES "peliculas"("id") ON DELETE CASCADE`).
		IfNotExists().
		Exec(context.TODO())
}
