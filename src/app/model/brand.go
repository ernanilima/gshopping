package model

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"created_date"`
}
