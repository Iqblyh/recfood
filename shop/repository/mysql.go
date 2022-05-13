package repository

import (
	"github.com/Iqblyh/recfood/shop/domain"
	"gorm.io/gorm"
)

type shopRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (shopRepo) Delete(id int) error {
	panic("unimplemented")
}

// GetById implements domain.Repository
func (shopRepo) GetById(id int) (domain.Shop, error) {
	panic("unimplemented")
}

// GetShops implements domain.Repository
func (shopRepo) GetShops() ([]domain.Shop, error) {
	panic("unimplemented")
}

// Save implements domain.Repository
func (shopRepo) Save(domain domain.Shop) (int, error) {
	panic("unimplemented")
}

// Update implements domain.Repository
func (shopRepo) Update(id int, domain domain.Shop) (domain.Shop, error) {
	panic("unimplemented")
}

func NewShopRepository(db *gorm.DB) domain.Repository {
	return shopRepo{
		DB: db,
	}
}
