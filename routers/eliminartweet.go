package routers

import (
	"redSGo/bd"
	"redSGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func EliminoTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {
	var r models.RespAPI
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) > 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio error al intentar borrar el tweet" + err.Error()
		return r
	}
	r.Message = "Tweet borrado con exito"
	r.Status = 200
	return r
}
