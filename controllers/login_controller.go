package controllers

import (
	"net/http"
	"os"
	"time"
	"warlock-backend/config"
	"warlock-backend/models"
	"warlock-backend/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var user models.User
		if err := config.DB.Preload("Roles").Where("email = ?", input.Email).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
			return
		}

		if user.IsBanned {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "User is banned"})
			return
		}

		if !util.CheckPassword(user.Password, input.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}

		roleNames := make([]models.RoleName, len(user.Roles))
		for i, role := range user.Roles {
			roleNames[i] = role.Name
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &models.Claims{
			Email:  user.Email,
			Roles:  roleNames,
			UserID: user.ID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
