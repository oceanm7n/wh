package auth

type UseCase interface {
	// Authenticate user on login, return token on success
	Login(username, password string) (string, error)
	// Register user on register, return token on success
	Register(username, password string) (string, error)
	// Parse token
	ParseToken(token string) (string, error)
}
