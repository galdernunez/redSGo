package routers

import (
	"encoding/json"
	"redSGo/bd"
	"redSGo/models"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespAPI {
	var r models.RespAPI
	r.Status = 400
	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "el parametro ID es obligatorio"
		return r
	}

	if len(pagina) < 1 {
		pagina = "1"
	}
	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "pagina debe ser un numerico mayor que 0"
		return r
	}
	tweets, correcto := bd.LeoTweets(ID, int64(pag))

	if !correcto {
		r.Message = "error al leer los tweets"
		return r
	}
	respJson, err := json.Marshal((tweets))

	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
