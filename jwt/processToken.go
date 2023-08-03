package jwt

import (
	"errors"
	"redSGo/models"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("invalid format Token")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {

	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("invalid Token")
	}

	return &claims, false, string(""), errors.New("error process Token")
}
