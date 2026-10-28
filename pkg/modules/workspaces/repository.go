package workspaces

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	DB *gorm.DB
}

func InitWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		DB: db,
	}
}

func (r *WorkspaceRepository) FindByID(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	var workspaceFinded *Workspace

	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&workspaceFinded).Error; err != nil {
		return nil, err
	}

	return workspaceFinded, nil
}

func (r *WorkspaceRepository) Create(ctx context.Context, entity *Workspace) (*Workspace, error) {
	if err := r.DB.WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *WorkspaceRepository) Update(ctx context.Context, id uuid.UUID, updates map[string]any) (*Workspace, error) {
	workspace, err := r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(workspace).Updates(updates).Error; err != nil {
		return nil, err
	}

	return workspace, nil
}

func (r *WorkspaceRepository) Delete(ctx context.Context, id uuid.UUID) error {

	result := r.DB.WithContext(ctx).Delete(&Workspace{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
