package users

import (
	"time"

	"github.com/google/uuid"
)

type ReadUserResponse struct {
	Id        *uuid.UUID `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Score     float32    `json:"score"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
