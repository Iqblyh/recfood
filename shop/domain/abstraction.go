package domain

type Service interface {
	GetDatas() ([]Shop, error)
	GetDataByID(id int) (Shop, error)
	CreateData(domain Shop) (Shop, error)
	EditData(id int) (Shop, error)
	DeleteData(id int) error
}

type Repository interface {
	GetShops() ([]Shop, error)
	GetById(id int) (Shop, error)
	Save(domain Shop) (int, error)
	Update(id int, domain Shop) (Shop, error)
	Delete(id int) error
}
