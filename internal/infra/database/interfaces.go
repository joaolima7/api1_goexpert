package database

import "github.com/joaolima7/api1_goexpert/internal/entity"

type UserInterface interface {
	Create(user entity.User) error
	FindByEmail(email string) (entity.User, error)
}

type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
