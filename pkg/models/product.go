package models

import (
	"time"
)

// Product represents a product.
// It contains the following fields:
// - ID: the unique identifier of the product
// - Name: the name of the product
// - Description: the description of the product
// - CreatedAt: the time when the product was created
type Product struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name" gorm:"not null" binding:"required"`
	Description string    `json:"description" gorm:"not null" binding:"required"`
	CreatedAt   time.Time `json:"created_at" time_format:"2006-01-02"`
}
