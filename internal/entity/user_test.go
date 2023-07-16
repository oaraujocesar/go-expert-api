package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@doe.com", "123456")
	assert.NoError(t, err)

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "john@doe.com", user.Email)
	assert.Equal(t, "John Doe", user.Name)
}

func TestNewUserWithInvalidPassword(t *testing.T) {
	user, err := NewUser("John Doe", "john@doe.com", "")
	fmt.Println(user)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidEntity, err)
}

func TestValidatePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "john@doe.com", "123456")

	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
	assert.NotEqual(t, user.Password, "123456")
}
