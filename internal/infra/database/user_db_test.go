package database

import (
	"testing"

	"github.com/joaolima7/api1_goexpert/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := OpenTestDB()
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jhon", "t@gmail.com", "113421")
	userDB := NewUser(db)

	err = userDB.Create(*user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Password, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := OpenTestDB()
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jhon", "t@gmail.com", "113421")
	userDB := NewUser(db)

	err = userDB.Create(*user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Password, userFound.Password)
}
