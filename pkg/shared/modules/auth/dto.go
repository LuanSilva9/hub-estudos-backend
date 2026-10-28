package auth

import "github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"

type RegisterResponse struct {
	Message string                  `json:"message"`
	User    *users.ReadUserResponse `json:"user"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresAt int    `json:"expires_at"`
}
