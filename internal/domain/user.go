package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id"          db:"id"`
	Login         string    `json:"login"       db:"login"`
	Email         string    `json:"email"       db:"email"`
	Password      string    `json:"password"    db:"password"`
	CreatedAt     time.Time `json:"created_at"  db:"created_at"`
	ProfileImage  string    `json:"profile_image" db:"profile_image"`
	EmailVerified bool      `json:"email_verified" db:"email_verified"`
}
