package routers

import (
	"context"
	"encoding/json"
	"redSGo/bd"
	"redSGo/models"
	"time"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespAPI {
	var mensaje models.Tweet
	var r models.RespAPI
	r.Status = 400
	IDUser := claim.ID.Hex()
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil {
		r.Message = "Ocurrio error al decodificar el body: " + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserId:  IDUser,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		r.Message = "Ocurrio error al insertar el registro: " + err.Error()
		r.Status = 500
		return r
	}
	if !status {
		r.Message = "No se ha logrado insertar el registro: "
		r.Status = 500
		return r
	}

	r.Message = "tweet Insertado"
	r.Status = 200
	return r
}
