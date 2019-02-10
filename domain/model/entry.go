package model

import (
	"time"
)

type Entry struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"not null" json:"title" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
