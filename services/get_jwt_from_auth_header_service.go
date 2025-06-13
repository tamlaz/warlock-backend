package services

import (
	"errors"
	"os"
	"strings"
	"warlock-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func GetClaimsFromAuthHeader(ctx *gin.Context) (*models.Claims, error) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		err := errors.New("authorization header is missing from request")
		return nil, err
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == "" {
		err := errors.New("JWT token is missing in Authorization header")
		return nil, err
	}

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
