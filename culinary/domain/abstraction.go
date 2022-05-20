package domain

type Service interface {
	GetDatas() (response []Culinary, err error)
	GetDataByID(id int) (response Culinary, err error)
	GetDataByCat(cat string) (response []Culinary, err error)
	GetDataByShopID(shop_id int) (response []Culinary, err error)
	GetByName(query string) (response []Culinary, err error)
	GetRecomendation(weather_cat string) (response Culinary, err error)
	InsertData(shop_id int, domain Culinary) (response Culinary, err error)
	EditData(id int, domain Culinary) (response Culinary, err error)
	DeleteData(id int) error
	CheckTemp() (string, error)
}

type Repository interface {
	GetCulinarys() ([]Culinary, error)
	GetById(id int) (Culinary, error)
	GetByShopID(shop_id int) ([]Culinary, error)
	GetByCat(cat string) ([]Culinary, int64)
	GetByWeatherCat(weather_cat string) (Culinary, error)
	GetTemp() (Weather, error)
	Search(query string) ([]Culinary, error)
	Save(shop_id int, domain Culinary) (int, error)
	Update(id int, domain Culinary) error
	Delete(id int) error
}
