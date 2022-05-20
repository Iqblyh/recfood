package repository

import (
	"time"

	"github.com/Iqblyh/recfood/review/domain"
	"gorm.io/gorm"
)

type reviewRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (rr reviewRepo) Delete(id int) error {
	var rec Review
	err := rr.DB.Unscoped().Delete(&rec, id).Error

	return err
}

// GetById implements domain.Repository
func (rr reviewRepo) GetById(id int) (domain.Review, error) {
	var newRecord Review
	err := rr.DB.Where("id = ?", id).First(&newRecord).Error

	return toDomain(newRecord), err
}

// GetReviews implements domain.Repository
func (rr reviewRepo) GetReviews(culinary_id int) ([]domain.Review, error) {
	records := []Review{}
	err := rr.DB.Where("culinary_id = ?", culinary_id).Find(&records).Error

	var culinarys []domain.Review
	for _, value := range records {
		culinarys = append(culinarys, toDomain(value))
	}

	return culinarys, err
}

// GetUserReviews implements domain.Repository
func (rr reviewRepo) GetUserReviews(user_id int) ([]domain.Review, error) {
	records := []Review{}
	err := rr.DB.Where("user_id = ?", user_id).Find(&records).Error

	var culinarys []domain.Review
	for _, value := range records {
		culinarys = append(culinarys, toDomain(value))
	}

	return culinarys, err
}

// Save implements domain.Repository
func (rr reviewRepo) Save(culinary_id, user_id int, domain domain.Review) (int, error) {
	newRecord := fromDomain(domain)
	newRecord.Culinary_Id = culinary_id
	newRecord.User_Id = user_id
	err := rr.DB.Save(&newRecord).Error

	return int(newRecord.ID), err
}

// Update implements domain.Repository
func (rr reviewRepo) Update(id int, domain domain.Review) error {
	newRecord := map[string]interface{}{
		"rating":     domain.Rating,
		"review":     domain.Review,
		"Updated_At": time.Now(),
	}

	err := rr.DB.Model(&domain).Where("id = ?", id).Updates(newRecord).Error

	return err
}

func NewReviewRepository(db *gorm.DB) domain.Repository {
	return reviewRepo{
		DB: db,
	}
}
