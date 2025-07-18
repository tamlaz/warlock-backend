package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Email  string
	Roles  []RoleName
	UserID uint
	jwt.RegisteredClaims
}
