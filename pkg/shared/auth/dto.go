package auth

type CredentialsRegister struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CredentialsLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
