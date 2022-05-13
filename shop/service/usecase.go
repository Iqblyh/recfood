package service

import "github.com/Iqblyh/recfood/shop/domain"

type shopService struct {
	repository domain.Repository
}

// CreateData implements domain.Service
func (shopService) CreateData(domain domain.Shop) (domain.Shop, error) {
	panic("unimplemented")
}

// DeleteData implements domain.Service
func (shopService) DeleteData(id int) error {
	panic("unimplemented")
}

// EditData implements domain.Service
func (shopService) EditData(id int) (domain.Shop, error) {
	panic("unimplemented")
}

// GetDataByID implements domain.Service
func (shopService) GetDataByID(id int) (domain.Shop, error) {
	panic("unimplemented")
}

// GetDatas implements domain.Service
func (shopService) GetDatas() ([]domain.Shop, error) {
	panic("unimplemented")
}

func NewShopService(repo domain.Repository) domain.Service {
	return shopService{
		repository: repo,
	}
}
