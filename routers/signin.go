package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"redSGo/bd"
	"redSGo/models"
)

func SignIn(ctx context.Context) models.RespAPI {
	var t models.User
	var r models.RespAPI

	r.Status = 400

	fmt.Println("IN SignIn")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "You must specify the email"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "You must specify a password of at least 6 characters"
		fmt.Println(r.Message)
		return r
	}

	_, found, _ := bd.ChekExistUser(t.Email)
	if found {
		r.Message = "You must specify a password of at least 6 characters"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		r.Message = "An error occurred while trying to register the user"
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "Failed to insert user record"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "SigIn OK"
	fmt.Println(r.Message)
	return r
}
