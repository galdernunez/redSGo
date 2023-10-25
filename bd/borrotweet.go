package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(ID string, UserID string) error {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	where := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, where)
	return err
}
