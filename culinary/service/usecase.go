package service

import (
	"errors"
	"fmt"

	"github.com/Iqblyh/recfood/culinary/domain"
	errorConv "github.com/Iqblyh/recfood/helper/error"
)

type culinaryService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (cs culinaryService) DeleteData(id int) error {
	err := cs.repository.Delete(id)
	if err != nil {
		return errorConv.Conversion(err)
	}
	return nil
}

// EditData implements domain.Service
func (cs culinaryService) EditData(id int, domain domain.Culinary) (response domain.Culinary, err error) {
	err = cs.repository.Update(id, domain)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	data, err := cs.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetByCat implements domain.Service
func (cs culinaryService) GetDataByCat(cat string) (response []domain.Culinary, err error) {
	data, rowAffected := cs.repository.GetByCat(cat)
	if rowAffected == 0 {
		err = errors.New("record not found")
		return response, err
	}
	return data, nil
}

// GetByName implements domain.Service
func (cs culinaryService) GetByName(query string) (response []domain.Culinary, err error) {
	data, err := cs.repository.Search(query)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetDataByID implements domain.Service
func (cs culinaryService) GetDataByID(id int) (response domain.Culinary, err error) {
	data, err := cs.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

//GetDataByShopID implements domain.Service
func (cs culinaryService) GetDataByShopID(shop_id int) (response []domain.Culinary, err error) {
	data, err := cs.repository.GetByShopID(shop_id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	return data, nil
}

// GetDatas implements domain.Service
func (cs culinaryService) GetDatas() (response []domain.Culinary, err error) {
	data, err := cs.repository.GetCulinarys()
	if err != nil {
		return response, errorConv.Conversion(err)
	}

	return data, nil
}

// GetRecomendation implements domain.Service
func (cr culinaryService) GetRecomendation(weather_cat string) (response domain.Culinary, err error) {
	data, err := cr.repository.GetByWeatherCat(weather_cat)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	fmt.Println(data)
	return data, nil
}

// InsertData implements domain.Service
func (cs culinaryService) InsertData(shop_id int, domain domain.Culinary) (response domain.Culinary, err error) {
	id, err := cs.repository.Save(shop_id, domain)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	data, err := cs.repository.GetById(id)
	if err != nil {
		return response, errorConv.Conversion(err)
	}
	return data, err
}

func (cs culinaryService) CheckTemp() (string, error) {
	data, err := cs.repository.GetTemp()

	if err != nil {
		return "wadoh", errorConv.Conversion(err)
	}

	temp := data.Main.Temp
	fmt.Println(temp)

	if temp < 298.15 {
		return "dingin", nil
	}
	return "panas", nil
}

func NewCulinaryService(repo domain.Repository) domain.Service {
	return culinaryService{
		repository: repo,
	}
}
