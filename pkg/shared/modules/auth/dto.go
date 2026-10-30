package auth

type RegisterResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresAt int    `json:"expires_at"`
}
