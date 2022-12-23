package entity

import (
	"time"
)

// Product ...
type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
