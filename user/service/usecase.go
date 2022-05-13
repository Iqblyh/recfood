package service

import (
	errorConv "github.com/Iqblyh/recfood/helper/error"
	"github.com/Iqblyh/recfood/user/domain"
)

type userService struct {
	repository domain.Repository
}

// CreateToken implements domain.Service
func (userService) CreateToken(email string, password string) (token string, err error) {
	panic("unimplemented")
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

func NewUserService(repo domain.Repository) domain.Service {
	return userService{
		repository: repo,
	}
}
