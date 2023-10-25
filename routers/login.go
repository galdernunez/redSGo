package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"redSGo/bd"
	"redSGo/jwt"
	"redSGo/models"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.RespAPI {
	var t models.User
	var r models.RespAPI
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		r.Message = "User/Password Invalid " + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "Email's user is mandatory"
		return r
	}

	userData, existe := bd.IntentoLogin(t.Email, t.Password)

	if !existe {
		r.Message = "User/Password Invalid"
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)

	if err != nil {
		r.Message = "Error al generar el token correspondiente " + err.Error()
		return r
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)

	if err2 != nil {
		r.Message = "Error al formatear el token correspondiente " + err2.Error()
		return r
	}

	cookies := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookies.String()
	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
