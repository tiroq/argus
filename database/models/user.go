package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int64     `gorm:"type:bigint;unique;not null"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Username  string    `gorm:"type:varchar(255);unique;not null"`
	JoinDate  time.Time `gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
	LastSeen  time.Time `gorm:"datetime:timestamp;default:CURRENT_TIMESTAMP"`
}
