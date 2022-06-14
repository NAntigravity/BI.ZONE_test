package models

import "github.com/dgrijalva/jwt-go"

type JWT struct {
	UserID uint
	Role   uint
	jwt.StandardClaims
}
