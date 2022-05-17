package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/Iqblyh/recfood/culinary/domain"
	"gorm.io/gorm"
)

type culinaryRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (cr culinaryRepo) Delete(id int) error {
	var rec Culinary
	err := cr.DB.Unscoped().Delete(&rec, id).Error

	return err
}

// GetByCat implements domain.Repository
func (cr culinaryRepo) GetByCat(cat string) ([]domain.Culinary, int64) {
	records := []Culinary{}
	err := cr.DB.Where("category = ?", cat).Find(&records).RowsAffected

	var culinarys []domain.Culinary

	for _, value := range records {
		culinarys = append(culinarys, toDomain(value))
	}
	return culinarys, err
}

// GetById implements domain.Repository
func (cr culinaryRepo) GetById(id int) (domain.Culinary, error) {
	var newRecord Culinary
	err := cr.DB.Where("id = ?", id).First(&newRecord).Error

	return toDomain(newRecord), err
}

// GetByShopID implements domain.Repository
func (cr culinaryRepo) GetByShopID(shop_id int) ([]domain.Culinary, error) {
	records := []Culinary{}
	err := cr.DB.Where("shop_id = ?", shop_id).Find(&records).Error

	var culinarys []domain.Culinary
	for _, value := range records {
		culinarys = append(culinarys, toDomain(value))
	}

	return culinarys, err
}

// GetCulinarys implements domain.Repository
func (cr culinaryRepo) GetCulinarys() ([]domain.Culinary, error) {
	records := []Culinary{}
	err := cr.DB.Find(&records).Error

	var culinarys []domain.Culinary

	for _, value := range records {
		culinarys = append(culinarys, toDomain(value))
	}
	return culinarys, err
}

// GetTemp implements domain.Repository
func (culinaryRepo) GetTemp() (domain.Weather, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?lat=-7.3961114&lon=109.69516&appid=e17feb8e52d2a152478637d7762b6a40"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	dataByte, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		fmt.Println(err.Error())
	}

	var respData domain.Weather
	json.Unmarshal(dataByte, &respData)
	// fmt.Println(body)

	return respData, nil
}

// Save implements domain.Repository
func (cr culinaryRepo) Save(shop_id int, domain domain.Culinary) (int, error) {
	newRecord := fromDomain(domain)
	newRecord.Shop_Id = shop_id
	err := cr.DB.Save(&newRecord).Error

	return int(newRecord.ID), err
}

// Search implements domain.Repository
func (cr culinaryRepo) Search(query string) ([]domain.Culinary, error) {
	records := []Culinary{}
	err := cr.DB.Where("culinary_name LIKE ?", "%"+query+"%").Find(&records).Error

	var culinarys []domain.Culinary
	for _, value := range records {
		culinarys = append(culinarys, toDomain(value))
	}
	return culinarys, err
}

// Update implements domain.Repository
func (cr culinaryRepo) Update(id int, domain domain.Culinary) error {
	newRecord := map[string]interface{}{
		"category":      domain.Category,
		"culinary_name": domain.Culinary_Name,
		"price":         domain.Price,
		"description":   domain.Description,
		"weather_cat":   domain.Weather_cat,
	}

	err := cr.DB.Model(&domain).Where("id = ?", id).Updates(newRecord).Error

	return err
}

func (cr culinaryRepo) GetByWeatherCat(weather_cat string) (domain.Culinary, error) {
	records := []Culinary{}
	err := cr.DB.Where("weather_cat = ?", weather_cat).Find(&records).Error
	count := len(records)
	var newRecords []domain.Culinary

	for _, value := range records {
		newRecords = append(newRecords, toDomain(value))
	}
	index := rand.Intn(count)
	return newRecords[index], err
}

func NewCulinaryRepository(db *gorm.DB) domain.Repository {
	return culinaryRepo{
		DB: db,
	}
}
