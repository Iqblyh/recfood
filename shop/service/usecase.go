package service

import (
	errorConv "github.com/Iqblyh/recfood/helper/error"
	"github.com/Iqblyh/recfood/shop/domain"
)

type shopService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (ss shopService) DeleteData(id int) error {
	err := ss.repository.Delete(id)
	if err != nil {
		return errorConv.Conversion(err)
	}
	return nil
}

// EditData implements domain.Service
func (ss shopService) EditData(id int, domain domain.Shop) (response domain.Shop, err error) {
	err = ss.repository.Update(id, domain)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	data, err := ss.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetDataByID implements domain.Service
func (ss shopService) GetDataByID(id int) (response domain.Shop, err error) {
	data, err := ss.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetDatas implements domain.Service
func (ss shopService) GetDatas() (response []domain.Shop, err error) {
	data, err := ss.repository.GetShops()
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// InsertData implements domain.Service
func (ss shopService) InsertData(domain domain.Shop) (response domain.Shop, err error) {
	id, err := ss.repository.Save(domain)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	data, err := ss.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	return data, err
}

func NewShopService(repo domain.Repository) domain.Service {
	return shopService{
		repository: repo,
	}
}
