package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           *uuid.UUID `gorm:"type:UUID;primaryKey;default:gen_random_uuid()"`
	Username     string     `gorm:"type:VARCHAR(255);not null"`
	Email        string     `gorm:"type:VARCHAR(255);not null;uniqueIndex"`
	PasswordHash string     `gorm:"type:VARCHAR(255);not null"`
	Score        float32    `gorm:"type:NUMERIC(10, 2);default:0"`
	CreatedAt    *time.Time `gorm:"type:TIMESTAMP;default:NOW()"`
	UpdatedAt    *time.Time `gorm:"type:TIMESTAMP"`
}

func (User) TableName() string {
	return "users"
}
