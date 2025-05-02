package database

import "github.com/joaolima7/api1_goexpert/internal/entity"

type UserInterface interface {
	Create(user entity.User) error
	FindByEmail(email string) (entity.User, error)
}
