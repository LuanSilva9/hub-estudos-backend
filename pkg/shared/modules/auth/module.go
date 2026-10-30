package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthModule(r *gin.RouterGroup, db *gorm.DB) {
	authHandler := InitAuthHandler(db)

	r.POST("/auth/signin", authHandler.Login)
	r.POST("/auth/register", authHandler.Register)
}
