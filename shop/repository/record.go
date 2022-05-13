package repository

import (
	"time"

	"github.com/Iqblyh/recfood/shop/domain"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Shop_Name string
	Address   string
	Open      string
	Close     string
}

func toDomain(rec Shop) domain.Shop {
	return domain.Shop{
		Shop_Name: rec.Shop_Name,
		Address:   rec.Address,
		Open:      rec.Open,
		Close:     rec.Close,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func fromDomain(rec domain.Shop) Shop {
	return Shop{
		Shop_Name: rec.Shop_Name,
		Address:   rec.Address,
		Open:      rec.Open,
		Close:     rec.Close,
	}
}
