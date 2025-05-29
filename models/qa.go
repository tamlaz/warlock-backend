package models

import "time"

type Qa struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"one2one"`
	Question  string `gorm:"notnull"`
	Answer    string `gorm:"notnull"`
	CreatedAt time.Time
}
