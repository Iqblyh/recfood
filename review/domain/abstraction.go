package domain

type Service interface {
	GetDatas(culinary_id int) (response []Review, err error)
	GetUserDatas(user_id int) (response []Review, err error)
	GetDataByID(id int) (response Review, err error)
	InsertData(culinary_id, user_id int, domain Review) (response Review, err error)
	EditData(id int, domain Review) (response Review, err error)
	DeleteData(id int) error
}

type Repository interface {
	GetReviews(culinary_id int) ([]Review, error)
	GetById(id int) (Review, error)
	GetUserReviews(user_id int) ([]Review, error)
	Save(culinary_id, user_id int, domain Review) (int, error)
	Update(id int, domain Review) error
	Delete(id int) error
}
