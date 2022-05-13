package repository

import (
	"time"

	"github.com/Iqblyh/recfood/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Fullname string
}

func toDomain(rec User) domain.User {
	return domain.User{
		Username: rec.Username,
		Email:    rec.Email,
		Password: rec.Password,
		Fullname: rec.Fullname,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func fromDomain(rec domain.User) User {
	return User{
		Username: rec.Username,
		Email:    rec.Email,
		Password: rec.Password,
		Fullname: rec.Fullname,
	}
}
