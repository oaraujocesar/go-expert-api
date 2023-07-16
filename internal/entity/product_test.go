package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Empty(t, product.UpdatedAt)
	assert.Empty(t, product.DeletedAt)
}

func TestNewProduct_InvalidName(t *testing.T) {
	product, err := NewProduct("", 10.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)

	product, err = NewProduct("Product 1", -10.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	product, _ := NewProduct("Product 1", 10.0)

	assert.Nil(t, product.Validate())
}
