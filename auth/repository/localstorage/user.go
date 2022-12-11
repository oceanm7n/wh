package localstorage

import (
	"context"
	"errors"
	"sync"

	"github.com/oceanm7n/wh/models"
)

type UserLocalStorage struct {
	users map[string]*models.User
	mutex sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*models.User),
		mutex: sync.Mutex{},
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.users[user.ID] = user

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
