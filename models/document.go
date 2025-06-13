package models

import "time"

type Document struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	FileName     string `gorm:"notnull"`
	FilePath     string `gorm:"notnull"`
	SubjectId    uint   `gorm:"notnull"`
	TopicId      uint   `gorm:"notnull"`
	DocumentType string `gorm:"notnull"`
	CreatedAt    time.Time
}
