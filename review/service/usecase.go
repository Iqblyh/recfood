package service

import (
	errorConv "github.com/Iqblyh/recfood/helper/error"
	"github.com/Iqblyh/recfood/review/domain"
)

type reviewService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (rs reviewService) DeleteData(id int) error {
	err := rs.repository.Delete(id)
	if err != nil {
		return errorConv.Conversion(err)
	}
	return nil
}

// EditData implements domain.Service
func (rs reviewService) EditData(id int, domain domain.Review) (response domain.Review, err error) {
	err = rs.repository.Update(id, domain)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	data, err := rs.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetDataByID implements domain.Service
func (rs reviewService) GetDataByID(id int) (response domain.Review, err error) {
	data, err := rs.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetDatas implements domain.Service
func (rs reviewService) GetDatas(culinary_id int) (response []domain.Review, err error) {
	data, err := rs.repository.GetReviews(culinary_id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	return data, nil
}

// GetUserDatas implements domain.Service
func (rs reviewService) GetUserDatas(user_id int) (response []domain.Review, err error) {
	data, err := rs.repository.GetUserReviews(user_id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	return data, nil
}

// InsertData implements domain.Service
func (rs reviewService) InsertData(culinary_id, user_id int, domain domain.Review) (response domain.Review, err error) {
	id, err := rs.repository.Save(culinary_id, user_id, domain)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	data, err := rs.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	return data, err
}

func NewReviewService(repo domain.Repository) domain.Service {
	return reviewService{
		repository: repo,
	}
}
