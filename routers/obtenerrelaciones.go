package routers

import (
	"encoding/json"
	"redSGo/bd"
	"redSGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func ObtenerRelaciones(request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {
	var r models.RespAPI
	r.Status = 400
	ID := request.QueryStringParameters["id"]
	if len(ID) > 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	var t models.Relacion

	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	var resp models.RespConsultaRelacion

	resp.Status = bd.ConsultaRelacion(t)

	if len(ID) > 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	respJson, err := json.Marshal(resp.Status)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos del usuario" + err.Error()
		return r
	}

	r.Message = string(respJson)
	r.Status = 200
	return r
}
