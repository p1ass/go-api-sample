package model

import (
	"time"
)

type Tag struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"unique;not null" json:"name" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
