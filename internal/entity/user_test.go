package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Joao", "test@gmail.com", "123456")
	if err != nil {
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, user.ID)
		assert.NotEmpty(t, user.Password)
		assert.Equal(t, "Joao", user.Name)
		assert.Equal(t, "test@gmail.com", user.Email)
	}
}

func TestUser_ValidatorPassword(t *testing.T) {
	user, err := NewUser("Joao", "test@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePasword("123456"))
	assert.False(t, user.ValidatePasword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
