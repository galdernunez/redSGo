package bd

import (
	"context"
	"redSGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ConsultaRelacion(t models.Relacion) bool {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("relacion")

	where := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	err := col.FindOne(ctx, where).Decode(resultado)
	return err == nil
}
