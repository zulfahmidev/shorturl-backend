package dto

import (
	"time"

	"github.com/google/uuid"
)

type URLData struct {
	ID        uuid.UUID `json:"id" db:"id"`
	URL       string    `json:"url" db:"url"`
	Token     string    `json:"token" db:"token"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
