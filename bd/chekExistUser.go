package bd

import (
	"context"
	"redSGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ChekExistUser(email string) (models.User, bool, string) {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
