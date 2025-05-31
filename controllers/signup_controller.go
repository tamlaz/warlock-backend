package controllers

import (
	"net/http"
	"warlock-backend/config"
	"warlock-backend/models"
	"warlock-backend/util"

	"github.com/gin-gonic/gin"
)

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input struct {
			Email     string `json:"email"`
			Password  string `json:"password"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var existingUser models.User
		if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is an active registration with this email"})
			return
		}

		hashedPassword, _ := util.HashPassword(input.Password)

		var studentRole models.Role
		if err := config.DB.First(&studentRole, "name = ?", models.Student).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Student role not found in database"})
			return
		}
		newUser := models.User{
			Email:     input.Email,
			Password:  hashedPassword,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Roles:     []models.Role{studentRole},
			IsBanned:  false,
			Strikes:   0,
		}

		if err := config.DB.Create(&newUser).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}
