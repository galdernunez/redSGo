package routers

import (
	"redSGo/bd"
	"redSGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func EliminoRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {
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

	status, err := bd.DeleteRelacion(t)

	if err != nil {
		r.Status = 500
		r.Message = "Ocurrio un error al intentar borrar relacion: " + err.Error()
		return r
	}

	if !status {
		r.Status = 500
		r.Message = "No se borro la relacion: " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Relacion borrada"
	return r
}
