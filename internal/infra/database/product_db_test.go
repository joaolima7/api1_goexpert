package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joaolima7/api1_goexpert/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := OpenTestDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Notebook", 1200.00)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	productDb := NewProduct(db)
	err = productDb.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotEqual(t, product.ID, "")
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := OpenTestDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 22", products[1].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := OpenTestDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Notebook", 1200.00)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	productDb := NewProduct(db)
	err = productDb.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotEqual(t, product.ID, "")
	assert.NotEmpty(t, product.ID)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
}

func TestUpdateProduct(t *testing.T) {
	db, err := OpenTestDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Notebook", 1200.00)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	productDb := NewProduct(db)
	err = productDb.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotEqual(t, product.ID, "")
	assert.NotEmpty(t, product.ID)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)

	productFound.Name = "Updated Product"
	err = productDb.Update(productFound)
	assert.NoError(t, err)

	productUpdated, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Updated Product", productUpdated.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := OpenTestDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Notebook", 1200.00)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	productDb := NewProduct(db)
	err = productDb.CreateProduct(product)
	assert.NoError(t, err)
	assert.NotEqual(t, product.ID, "")
	assert.NotEmpty(t, product.ID)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)

	err = productDb.Delete(productFound.ID.String())
	assert.NoError(t, err)

	_, err = productDb.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
