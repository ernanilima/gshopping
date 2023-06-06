package model

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description" validate:"required,min=2,max=20"`
	CreatedAt   time.Time `json:"created_at"`
}
