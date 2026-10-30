package users

import (
	"net/http"

	"github.com/LuanSilva9/hub-estudos-backend/pkg/shared/exceptions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	repository *UserRepository
}

func InitUserHandler(repository *UserRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

func (h *UserHandler) FindMe(c *gin.Context) {
	userId, ok := c.MustGet("user-id").(string)

	userIdParsed, err := uuid.Parse(userId)

	if !ok || err != nil {
		exceptions.Unauthorized(c, "ID inválido")
		return
	}

	me, err := h.repository.FindByID(c.Request.Context(), userIdParsed)

	if err != nil {
		exceptions.NotFound(c, "Usuário não encontrado")
		return
	}

	c.JSON(http.StatusOK, EntityToDto(*me))
}
