package routers

import (
	"encoding/json"
	"fmt"
	"redSGo/bd"
	"redSGo/models"

	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespAPI {
	var r models.RespAPI
	r.Status = 400
	fmt.Println("Entré en VerPerfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) > 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrió un erro al intentar bucar el registro:" + err.Error()
		return r
	}
	respJson, err := json.Marshal((perfil))
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuario a JSON:" + err.Error()
		return r
	}
	r.Status = 500
	r.Message = string(respJson)
	return r
}
