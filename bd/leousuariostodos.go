package bd

import (
	"context"
	"fmt"
	"redSGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx := context.TODO()
	db := MongoCnn.Database(DatabaseName)
	col := db.Collection("usuarios")

	var result []*models.User

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSkip((page - 1) * 20)

	where := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, where, opciones)
	if err != nil {
		return result, false
	}

	for cur.Next(ctx) {
		var s models.User

		err := cur.Decode(&s)
		if err != nil {
			fmt.Println("decode: " + err.Error())
			return result, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir := false

		encontrado := false

		if tipo == "new" && !encontrado {
			incluir = true
		}

		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			result = append(result, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println("cru.Err(): " + err.Error())
		return result, false
	}
	cur.Close(ctx)
	return result, true
}
