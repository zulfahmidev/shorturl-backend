package model

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;default:gen_random_uuid()"`
	URL          string    `json:"url" gorm:"not null"`
	Token        string    `json:"token" gorm:"not null;unique;max=20"`
	VisitorCount int       `json:"visitor_count" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
