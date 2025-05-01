package entity

import (
	entit "github.com/joaolima7/api1_goexpert/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entit.ID `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
}

func NewUser(name string, email string, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entit.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePasword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
