package domain

type Service interface {
	GetDatas() (response []Shop, err error)
	GetDataByID(id int) (response Shop, err error)
	InsertData(domain Shop) (response Shop, err error)
	EditData(id int, domain Shop) (response Shop, err error)
	DeleteData(id int) error
}

type Repository interface {
	GetShops() ([]Shop, error)
	GetById(id int) (Shop, error)
	Save(domain Shop) (int, error)
	Update(id int, domain Shop) error
	Delete(id int) error
}
