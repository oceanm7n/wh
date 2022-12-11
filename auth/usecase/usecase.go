package usecase

import (
	"context"
	"crypto/sha256"
	"time"

	"github.com/oceanm7n/wh/auth"
	"github.com/oceanm7n/wh/models"
)

type AuthUseCaseV1 struct {
	userRepo       auth.Repository
	hashSalt       string
	expireDuration time.Duration
}

func NewAuthUseCasev1(
	userRepo auth.Repository,
	hashSalt string,
	expireDuration time.Duration,
) *AuthUseCaseV1 {
	return &AuthUseCaseV1{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		expireDuration: expireDuration,
	}
}

func (u *AuthUseCaseV1) CreateUser(ctx context.Context, user *models.User) error {
	// take the password and hash it
	// use sha256
	hash := sha256.Sum256([]byte(user.Password + u.hashSalt))
}
