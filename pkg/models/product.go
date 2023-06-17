package models

import (
	"time"
)

type Product struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" gorm:"not null" binding:"required"`
	Description string    `json:"description" gorm:"not null" binding:"required"`
	CreatedAt   time.Time `json:"created_at" time_format:"2006-01-02"`
}
