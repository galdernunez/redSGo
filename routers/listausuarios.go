package routers

import (
	"encoding/json"
	"redSGo/bd"
	"redSGo/models"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func ListaUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {
	var r models.RespAPI
	r.Status = 400
	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUser := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pagAux, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar el parametro 'page' como entero mayor a 0" + err.Error()
		return r
	}

	usuarios, status := bd.LeoUsuariosTodos(IDUser, int64(pagAux), search, typeUser)

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios" + err.Error()
		return r
	}

	if !status {
		r.Status = 500
		r.Message = "Error listando usuarios" + err.Error()
		return r
	}

	r.Message = string(respJson)
	r.Status = 200
	return r
}
