package model

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
