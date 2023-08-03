package bd

import (
	"context"
	"redSGo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u models.User) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("users")

	u.Password, _ = encodePassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
