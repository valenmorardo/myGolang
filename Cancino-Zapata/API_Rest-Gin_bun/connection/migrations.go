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
}
