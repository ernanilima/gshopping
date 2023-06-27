package model

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID            uuid.UUID `json:"id"`
	Code          int64     `json:"code"`
	Description   string    `json:"description" validate:"required,min=2,max=20"`
	TotalProducts int64     `json:"total_products"`
	CreatedAt     time.Time `json:"created_at"`
}
