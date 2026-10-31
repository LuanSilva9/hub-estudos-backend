package workspaces

import (
	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitWorkspaceModule(r *gin.RouterGroup, db *gorm.DB, middleware gin.HandlerFunc) {
	repository := InitWorkspaceRepository(db)
	userRepository := users.InitUserRepository(db)

	handler := InitWorkspaceHandler(repository, userRepository)

	r.POST("/new", middleware, handler.Create)
}
