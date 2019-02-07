package model

import (
	"time"
)

type Entry struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Title     string     `json:"title" validate:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
