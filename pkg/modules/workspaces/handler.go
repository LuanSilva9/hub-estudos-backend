package workspaces

import (
	"net/http"

	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkspaceHandler struct {
	repository     *WorkspaceRepository
	userRepository *users.UserRepository
}

func InitWorkspaceHandler(repository *WorkspaceRepository) *WorkspaceHandler {
	return &WorkspaceHandler{
		repository: repository,
	}
}

func (h *WorkspaceHandler) Create(c *gin.Context) {
	userId, ok := c.MustGet("user-id").(string)
	userIdParsed, err := uuid.Parse(userId)

	if !ok || err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário inválido"})
		return
	}

	var workspaceRequest CreateWorkspaceRequest

	if err := c.ShouldBindJSON(&workspaceRequest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao buscar usuário"})
		return
	}

	_, err = h.repository.Create(c.Request.Context(), &Workspace{
		Name:   workspaceRequest.Name,
		UserID: userIdParsed,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno ao criar usuário"})
		return
	}

}
