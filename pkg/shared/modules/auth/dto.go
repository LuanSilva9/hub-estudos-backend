package auth

import "github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"

type CredentialsRegisterRequest struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CredentialsLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string                  `json:"message"`
	User    *users.ReadUserResponse `json:"user"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresAt int    `json:"expires_at"`
}
