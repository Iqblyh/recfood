package repository

import (
	"github.com/Iqblyh/recfood/review/domain"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Culinary_Id int
	User_Id     int
	Rating      int
	Review      string
}

func toDomain(rec Review) domain.Review {
	return domain.Review{
		Culinary_Id: rec.Culinary_Id,
		User_Id:     rec.User_Id,
		Rating:      rec.Rating,
		Review:      rec.Review,
	}
}

func fromDomain(rec domain.Review) Review {
	return Review{
		Culinary_Id: rec.Culinary_Id,
		User_Id:     rec.User_Id,
		Rating:      rec.Rating,
		Review:      rec.Review,
	}
}
