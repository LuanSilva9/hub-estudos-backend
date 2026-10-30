package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserModule(r *gin.RouterGroup, db *gorm.DB, middleware gin.HandlerFunc) {
	userRepository := InitUserRepository(db)
	userHandler := InitUserHandler(userRepository)

	r.GET("/me", middleware, userHandler.FindMe)
}
