package model

import (
	"time"
)

type Tag struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
