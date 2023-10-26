package bd

import (
	"context"
	"redSGo/models"

	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int64) ([]models.RespTweetsSeguidores, bool) {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	where := make([]bson.M, 0)
	where = append(where, bson.M{"$match": bson.M{"usuarioid": ID}})
	where = append(where, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet"}})
	where = append(where, bson.M{"$unwind": "$tweet"})
	where = append(where, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	where = append(where, bson.M{"$skip": skip})
	where = append(where, bson.M{"$limit": 20})

	var result []models.RespTweetsSeguidores

	cursor, err := col.Aggregate(ctx, where)

	if err != nil {
		return result, false
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
