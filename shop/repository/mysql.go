package repository

import (
	"github.com/Iqblyh/recfood/shop/domain"
	"gorm.io/gorm"
)

type shopRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (sr shopRepo) Delete(id int) error {
	return sr.DB.Delete("id = ?", id).Error
}

// GetById implements domain.Repository
func (sr shopRepo) GetById(id int) (domain.Shop, error) {
	var newRecord Shop
	err := sr.DB.Where("id = ?", id).First(&newRecord).Error

	return toDomain(newRecord), err
}

// GetShops implements domain.Repository
func (sr shopRepo) GetShops() ([]domain.Shop, error) {
	records := []Shop{}
	err := sr.DB.Find(&records).Error

	var shops []domain.Shop

	for _, value := range records {
		shops = append(shops, toDomain(value))
	}
	return shops, err
}

// Save implements domain.Repository
func (sr shopRepo) Save(domain domain.Shop) (int, error) {
	newRecord := fromDomain(domain)
	err := sr.DB.Save(&newRecord).Error

	return int(newRecord.ID), err
}

// Update implements domain.Repository
func (sr shopRepo) Update(id int, domain domain.Shop) error {
	newRecord := map[string]interface{}{
		"shop_name": domain.Shop_Name,
		"address":   domain.Address,
		"Open":      domain.Open,
		"Close":     domain.Close,
	}

	err := sr.DB.Model(&domain).Where("id = ?", id).Updates(newRecord).Error

	return err
}

func NewShopRepository(db *gorm.DB) domain.Repository {
	return shopRepo{
		DB: db,
	}
}
