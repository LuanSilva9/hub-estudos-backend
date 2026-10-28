package workspaces

import (
	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/google/uuid"
)

type Workspace struct {
	Id     *uuid.UUID `gorm:"type:UUID;primaryKey;default:gen_random_uuid()"`
	Name   string     `gorm:"type:VARCHAR(255);not null"`
	UserID uuid.UUID
	User   users.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Workspace) TableName() string {
	return "workspaces"
}
