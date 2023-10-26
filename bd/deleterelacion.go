package bd

import (
	"context"
	"redSGo/models"
)

func DeleteRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
