package handlers

import (
	"context"
	"fmt"
	"redSGo/jwt"
	"redSGo/models"
	"redSGo/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("Procces " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("path")).(string))

	var r models.RespAPI
	r.Status = 400

	isOK, statusCode, msg, claim := validAuthorization(ctx, request)

	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "sigin":
			return routers.SignIn(ctx)
		case "login":
			return routers.Login(ctx)
		case "tweet":
			return routers.GraboTweet(ctx, claim)
		case "subirAvatar":
			return routers.UploadImage(ctx, "A", request, claim)
		case "subirBanner":
			return routers.UploadImage(ctx, "B", request, claim)
		case "altarelacion":
			return routers.AltaRelacion(ctx, request, claim)
		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verperfil":
			return routers.VerPerfil(request)
		case "leerTweets":
			return routers.LeoTweets(request)
		case "obtenerAvatar":
			return routers.ObtenerImagen(ctx, "A", request, claim)
		case "obtenerBanner":
			return routers.ObtenerImagen(ctx, "B", request, claim)
		case "consultarelacion":
			return routers.ObtenerRelaciones(request, claim)
		case "listarusuarios":
			return routers.ListaUsuarios(request, claim)
		case "leotweetsseguidores":
			return routers.LeoTweetsSeguidores(request, claim)

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "modificarperfil":
			return routers.ModificarPerfil(ctx, claim)
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "elimnartweet":
			return routers.EliminoTweet(request, claim)
		case "elimnarrelacion":
			return routers.EliminoRelacion(request, claim)
		}
	}
	r.Message = "Method invalid"
	return r
}

func validAuthorization(ctx context.Context, reques events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "sigin" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := reques.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Requiered Token", models.Claim{}
	}

	claim, status, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !status {
		if err != nil {
			fmt.Println("Error in Token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error in Token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}
	fmt.Println("Token OK")

	return true, 200, msg, *claim
}
