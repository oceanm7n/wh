package localstorage

import (
	"context"
	"testing"

	"github.com/oceanm7n/wh/models"
	"github.com/stretchr/testify/assert"
)

func TestLocalStorageAuth(t *testing.T) {

	// new local storage
	ls := NewUserLocalStorage()

	// create user
	user := &models.User{
		ID:       "1",
		Username: "test",
		Password: "test",
	}
	err := ls.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	// get user
	user, err = ls.GetUser(context.Background(), "test", "test")
	assert.NoError(t, err)
	assert.NotNil(t, user)

	// get user that does not exist
	user, err = ls.GetUser(context.Background(), "test", "test2")
	assert.Error(t, err)
	assert.Nil(t, user)
}
