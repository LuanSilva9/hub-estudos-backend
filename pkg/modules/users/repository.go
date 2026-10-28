package users

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	var userFinded User

	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&userFinded).Error; err != nil {
		return nil, err
	}

	return &userFinded, nil
}

func (r *UserRepository) FindByIDStr(ctx context.Context, id string) (*User, error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.FindByID(ctx, idParsed)
}
