package auth

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input users.CredentialsRegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Inválido"})
		return
	}

	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.ToLower(strings.TrimSpace((input.Email)))

	if input.Name == "" || len(input.Name) > 255 || input.Email == "" || len(input.Email) > 255 || len(input.Password) < 8 || len(input.Password) > 72 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro Interno, contate um administrador"})
		return
	}

	userCreated := users.User{
		Username:     input.Name,
		Email:        input.Email,
		PasswordHash: string(hash),
	}

	err = h.DB.WithContext(c.Request.Context()).Create(&userCreated).Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "E-mail cadastrado - Chave única violada!",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno",
		})

		return
	}

	c.JSON(http.StatusCreated, &RegisterResponse{
		Message: "Usuário Criado com Sucesso!",
		User:    users.EntityToDto(userCreated),
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input users.CredentialsLoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "JSON inválido",
		})

		return
	}

	email := strings.ToLower(strings.TrimSpace(input.Email))

	var user users.User

	err := h.DB.WithContext(c.Request.Context()).Where("email = ?", email).First(&user).Error

	if err != nil || bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(input.Password),
	) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "E-mail ou senha inválidos",
		})
		return
	}

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno",
		})
		return
	}

	ttlHours, err := strconv.Atoi(os.Getenv("JWT_TTL_HOURS"))

	if err != nil || ttlHours < 1 || ttlHours > 168 {
		ttlHours = 24
	}

	now := time.Now()

	claims := jwt.RegisteredClaims{
		Subject:  user.Id.String(),
		IssuedAt: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(
			now.Add(time.Duration(ttlHours) * time.Hour),
		),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno",
		})

		return
	}

	c.JSON(http.StatusOK, &AuthResponse{
		Token:     signed,
		TokenType: "Bearer",
		ExpiresAt: ttlHours * 24,
	})
}
