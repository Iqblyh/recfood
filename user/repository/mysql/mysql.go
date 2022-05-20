package repository

import (
	"fmt"

	"github.com/Iqblyh/recfood/user/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (ur userRepo) Delete(id int) (err error) {
	return ur.DB.Delete("id = ?", id).Error
}

// GetByEmailPassword implements domain.Repository
func (ur userRepo) GetByUsernamePassword(username string, password string) (domain domain.User, err error) {
	var record User
	err = ur.DB.Where("username = ? AND password = ?", username, password).First(&record).Error
	return toDomain(record), err
}

// GetById implements domain.Repository
func (ur userRepo) GetById(id int) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.Where("id = ?", id).First(&newRecord).Error

	fmt.Println(toDomain(newRecord))
	return toDomain(newRecord), err
}

// Save implements domain.Repository
func (ur userRepo) Save(domain domain.User) (id int, err error) {
	record := fromDomain(domain)
	err = ur.DB.Save(&record).Error

	fmt.Println(record.Id)

	return int(record.Id), err
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepo{
		DB: db,
	}
}
