package repository

import (
	"github.com/Iqblyh/recfood/culinary/domain"
	"gorm.io/gorm"
)

type Culinary struct {
	gorm.Model
	Shop_Id       int
	Category      string
	Culinary_Name string
	Price         int
	Description   string
	Weather_Cat   string
}

func toDomain(rec Culinary) domain.Culinary {
	return domain.Culinary{
		Shop_Id:       rec.Shop_Id,
		Category:      rec.Category,
		Culinary_Name: rec.Culinary_Name,
		Price:         rec.Price,
		Description:   rec.Description,
		Weather_cat:   rec.Weather_Cat,
	}
}

func fromDomain(rec domain.Culinary) Culinary {
	return Culinary{
		Shop_Id:       rec.Shop_Id,
		Category:      rec.Category,
		Culinary_Name: rec.Culinary_Name,
		Price:         rec.Price,
		Description:   rec.Description,
		Weather_Cat:   rec.Weather_cat,
	}
}
