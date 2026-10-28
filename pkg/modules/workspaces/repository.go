package workspaces

import "gorm.io/gorm"

type WorkspaceRepository struct {
	DB *gorm.DB
}

func InitWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		DB: db,
	}
}

func (r *WorkspaceRepository) Create(entity *Workspace) (*Workspace, error) {
	if err := r.DB.Create(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}
