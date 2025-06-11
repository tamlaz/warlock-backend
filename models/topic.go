package models

type Topic struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"notnull"`
	SubjectId uint   `gorm:"notnull"`
}
