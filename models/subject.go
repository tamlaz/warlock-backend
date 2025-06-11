package models

type Subject struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"notnull"`
	Topics []Topic `gorm:"foreignKey:SubjectId"`
}
