package bd

import (
	"context"
	"redSGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.User, error) {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("usuarios")

	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)
	where := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, where).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		return perfil, err
	}
	return perfil, nil
}
