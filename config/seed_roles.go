package config

import (
	"log"
	"warlock-backend/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: models.Student},
		{Name: models.Teacher},
		{Name: models.Favorite},
	}

	for _, role := range roles {
		var count int64
		db.Model(&models.Role{}).Where("name = ?", role.Name).Count(&count)
		if count == 0 {
			if err := db.Create(&role).Error; err != nil {
				log.Printf("Failed to populate DB with role %s: %v", role.Name, err)
			}
		}
	}
}
