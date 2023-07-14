package database

import (
	"github.com/oaraujocesar/go-expert-api/internal/entity"
	e "github.com/oaraujocesar/go-expert-api/pkg/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindById(id e.ID) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Delete(id e.ID) error {
	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID)
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		err = p.DB.Offset(offset).Limit(limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}
