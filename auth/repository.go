package auth

// Interface for a repository that can authenticate a user

import (
	"context"

	"github.com/oceanm7n/wh/models"
)

type Repository interface {
	// Authenticate user on login, return token on success
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
