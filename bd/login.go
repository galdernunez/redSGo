package bd

import (
	"redSGo/models"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := ChekExistUser(email)
	if !encontrado {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordByD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordByD, passwordBytes)

	if err != nil {
		return usu, false
	}
	return usu, true
}
