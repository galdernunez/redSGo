package routers

import (
	"context"
	"encoding/json"
	"redSGo/bd"
	"redSGo/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespAPI {
	var r models.RespAPI
	r.Status = 400

	var t models.User

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Datos Incorrecto: " + err.Error()
	}

	status, err := bd.ModificarPerfil(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al intentar modificar el registro: " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado modificar el registro: "
		return r
	}
	r.Status = 200
	r.Message = "Modificación de perfil OK"

	return r

}
