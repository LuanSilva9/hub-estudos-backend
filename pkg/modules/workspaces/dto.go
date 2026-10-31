package workspaces

import (
	"time"

	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/google/uuid"
)

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}

type ReadWorkspaceResponse struct {
	Id        *uuid.UUID             `json:"id"`
	Name      string                 `json:"name"`
	User      users.ReadUserResponse `json:"owner"`
	CreatedAt *time.Time             `json:"created_at"`
	UpdatedAt *time.Time             `json:"updated_at"`
}
