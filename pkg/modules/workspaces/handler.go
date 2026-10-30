package workspaces

import (
	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/LuanSilva9/hub-estudos-backend/pkg/shared/exceptions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkspaceHandler struct {
	repository     *WorkspaceRepository
	userRepository *users.UserRepository
}

func InitWorkspaceHandler(repository *WorkspaceRepository, userRepository *users.UserRepository) *WorkspaceHandler {
	return &WorkspaceHandler{
		repository:     repository,
		userRepository: userRepository,
	}
}

func (h *WorkspaceHandler) Create(c *gin.Context) {
	userId, ok := c.MustGet("user-id").(string)
	userIdParsed, err := uuid.Parse(userId)

	if !ok || err != nil {
		exceptions.Unauthorized(c, "Usuário Inválido ou sem autorização")
		return
	}

	var workspaceRequest CreateWorkspaceRequest

	if err := c.ShouldBindJSON(&workspaceRequest); err != nil {
		exceptions.BadRequest(c, "Requisição ou JSON inválidos")
		return
	}

	_, err = h.repository.Create(c.Request.Context(), &Workspace{
		Name:   workspaceRequest.Name,
		UserID: userIdParsed,
	})

	if err != nil {
		exceptions.InternalServerError(c, "Erro interno ao criar Workspace, contate um administrador")
		return
	}

}
