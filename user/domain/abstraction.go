package domain

type Service interface {
	Login(username, password string) (token string, err error)
	InsertData(domain User) (response User, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain User) (id int, err error)
	Delete(id int) (err error)
	GetByUsernamePassword(username, password string) (domain User, err error)
	GetById(id int) (domain User, err error)
}

type Weather interface {
	GetWeatherTemp(hmm, hma string) (temp float64)
}
