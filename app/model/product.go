package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Barcode     string    `json:"barcode"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductNotFound struct {
	ID       int64  `json:"id"`
	Barcode  string `json:"barcode"`
	Attempts string `json:"attempts"`
}
