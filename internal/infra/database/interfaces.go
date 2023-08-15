package database

import (
	"github.com/google/uuid"
	"github.com/oaraujocesar/go-expert-api/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id uuid.UUID) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id uuid.UUID) error
}
