package service

import (
	"errors"

	errorConv "github.com/Iqblyh/recfood/helper/error"
	"github.com/Iqblyh/recfood/middlewares"
	"github.com/Iqblyh/recfood/user/domain"
)

type userService struct {
	repository domain.Repository
	jwtAuth    *middlewares.ConfigJWT
}

// CreateToken implements domain.Service
func (us userService) Login(username string, password string) (token string, err error) {
	data, err := us.repository.GetByUsernamePassword(username, password)

	if err != nil {
		return "login gagal", errorConv.Conversion(err)
	}
	token, err = us.jwtAuth.GenerateToken(data.Id)

	if err != nil {
		return "gagal", errors.New("waduh")
	}
	return token, nil
}

// DeleteData implements domain.Service
func (us userService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)
	return errorConv.Conversion(errResp)
}

// InsertData implements domain.Service
func (us userService) InsertData(domain domain.User) (response domain.User, err error) {
	id, err := us.repository.Save(domain)

	if err != nil {
		return response, errorConv.Conversion(err)
	}

	res, errNew := us.repository.GetById(id)
	if errNew != nil {
		return response, errorConv.Conversion(errNew)
	}

	return res, nil

}

func NewUserService(repo domain.Repository, JWT *middlewares.ConfigJWT) domain.Service {
	return userService{
		repository: repo,
		jwtAuth:    JWT,
	}
}
