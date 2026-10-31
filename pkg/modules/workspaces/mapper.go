package workspaces

import "github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"

func EntityToDto(entity Workspace) *ReadWorkspaceResponse {
	return &ReadWorkspaceResponse{
		Id:        entity.Id,
		Name:      entity.Name,
		User:      *users.EntityToDto(entity.User),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
