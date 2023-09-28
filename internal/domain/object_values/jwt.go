package objectValues

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Id        uint64 `json:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Correo    string `json:"correo"`
	jwt.StandardClaims
}

type JwtEntity struct {
	Id        uint64 `json:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Correo    string `json:"correo"`
}
