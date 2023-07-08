package database

import (
	"testing"

	"github.com/oaraujocesar/go-expert-api/internal/entity"
	e "github.com/oaraujocesar/go-expert-api/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 13.0)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating an product entity", err)
	}

	productDB := NewProduct(db)
	productDB.Create(product)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	t.Run("It should get product by id", func(t *testing.T) {
		product, err := entity.NewProduct("Product 1", 13.0)
		assert.Nil(t, err)

		productDB := NewProduct(db)
		productDB.Create(product)

		productFound, err := productDB.FindById(product.ID)

		assert.Nil(t, err)
		assert.Equal(t, product.ID, productFound.ID)
		assert.Equal(t, product.Name, productFound.Name)
		assert.Equal(t, product.Price, productFound.Price)
	})

	t.Run("It should return error when product not found", func(t *testing.T) {
		productDB := NewProduct(db)

		_, err := productDB.FindById(e.NewID())

		assert.NotNil(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	t.Run("It should delete product by id", func(t *testing.T) {
		product, err := entity.NewProduct("Product 1", 13.0)
		assert.Nil(t, err)

		productDB := NewProduct(db)
		productDB.Create(product)

		err = productDB.Delete(product.ID)

		assert.Nil(t, err)

		var productFound entity.Product
		err = db.First(&productFound, "id = ?", product.ID).Error

		assert.NotNil(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	t.Run("It should update the price and name", func(t *testing.T) {
		product, err := entity.NewProduct("Product 1", 13.0)
		assert.Nil(t, err)

		productDB := NewProduct(db)
		productDB.Create(product)

		product.Price = 15.0
		product.Name = "Product 2"

		err = productDB.Update(product)

		assert.Nil(t, err)

		var productFound entity.Product
		err = db.First(&productFound, "id = ?", product.ID).Error

		assert.Nil(t, err)
		assert.Equal(t, product.ID, productFound.ID)
		assert.Equal(t, product.Name, productFound.Name)
		assert.Equal(t, product.Price, productFound.Price)

		assert.NotEqual(t, "Product 1", productFound.Name)
		assert.NotEqual(t, 13.0, productFound.Price)
	})

	t.Run("It should return error when product not found", func(t *testing.T) {
		product, err := entity.NewProduct("Product 1", 13.0)
		assert.Nil(t, err)

		productDB := NewProduct(db)

		err = productDB.Update(product)

		assert.NotNil(t, err)
	})
}
