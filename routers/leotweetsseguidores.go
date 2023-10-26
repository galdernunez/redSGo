package routers

import (
	"encoding/json"
	"redSGo/bd"
	"redSGo/models"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func LeoTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {
	var r models.RespAPI
	r.Status = 400
	IDUser := claim.ID.Hex()
	if len(IDUser) < 1 {
		r.Message = "el parametro ID es obligatorio"
		return r
	}
	pagina := request.QueryStringParameters["pagina"]
	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "pagina debe ser un numerico mayor que 0"
		return r
	}

	tweets, status := bd.LeoTweetsSeguidores(IDUser, int64(pag))

	if !status {
		r.Message = "Error al leer los tweets seguidores"
		return r
	}

	respJson, err := json.Marshal((tweets))

	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de tweet de los sequidores como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}
