package database

import (
	"testing"
	"time"

	"github.com/oaraujocesar/go-expert-api/internal/entity"
	e "github.com/oaraujocesar/go-expert-api/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createProducts(db *gorm.DB) {
	productDB := NewProduct(db)

	ps := []entity.Product{
		{Name: "P1", Price: 13.0},
		{Name: "P2", Price: 143.0},
		{Name: "P3", Price: 11.0},
		{Name: "P4", Price: 122.0},
	}

	for _, p := range ps {
		product, _ := entity.NewProduct(p.Name, p.Price)
		time.Sleep(time.Millisecond * 100)
		productDB.Create(product)
	}

}

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

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db.AutoMigrate(&entity.Product{})

	createProducts(db)

	t.Run("it should return all the products according with pagination and ascending order", func(t *testing.T) {
		productDB := NewProduct(db)

		products, err := productDB.FindAll(1, 2, "asc")
		assert.Nil(t, err)
		assert.Len(t, products, 2)
		assert.Greater(t, products[len(products)-1].CreatedAt, products[0].CreatedAt)
	})

	t.Run("it should return the products in ascending order when sort arg is empty", func(t *testing.T) {
		productDB := NewProduct(db)

		products, err := productDB.FindAll(1, 2, "")
		assert.Nil(t, err)
		assert.Len(t, products, 2)
		assert.Greater(t, products[len(products)-1].CreatedAt, products[0].CreatedAt)
	})

	t.Run("it should return the products in descending order when sort arg is empty", func(t *testing.T) {
		productDB := NewProduct(db)

		products, err := productDB.FindAll(1, 4, "desc")

		assert.Nil(t, err)
		assert.Len(t, products, 4)
		assert.Equal(t, products[0].Name, "P4")
	})

	t.Run("it should return an error if page and limit are 0", func(t *testing.T) {
		productDB := NewProduct(db)

		products, err := productDB.FindAll(0, 0, "")

		assert.Nil(t, err)
		assert.Len(t, products, 4)
		assert.Greater(t, products[len(products)-1].CreatedAt, products[0].CreatedAt)
	})
}
