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
