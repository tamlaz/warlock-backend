package models

import "time"

type Qa struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	User      User `gorm:"foreignKey:UserId"`
	SubjectId uint
	TopicId   uint
	Question  string `gorm:"notnull"`
	Answer    string `gorm:"notnull"`
	CreatedAt time.Time
}
