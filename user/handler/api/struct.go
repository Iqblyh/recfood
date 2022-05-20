package userApi

import "github.com/Iqblyh/recfood/user/domain"

type RequestJSON struct {
	Id       int
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Id:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Fullname: req.Fullname,
	}
}

type ResponseJSON struct {
	Id       int
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

func fromDomain(domain domain.User) ResponseJSON {
	return ResponseJSON{
		Id:       domain.Id,
		Username: domain.Username,
		Email:    domain.Email,
		Fullname: domain.Fullname,
	}
}
