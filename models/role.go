package models

type RoleName string

const (
	Student  RoleName = "STUDENT"
	Teacher  RoleName = "TEACHER"
	Favorite RoleName = "FAVORITE"
)

type Role struct {
	Id   uint     `gorm:"primaryKey"`
	Name RoleName `gorm:"not null"`
}
