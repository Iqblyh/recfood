package domain_test

import (
	"os"
	"testing"

	"github.com/Iqblyh/recfood/middlewares"
	userDomain "github.com/Iqblyh/recfood/user/domain"
	userMock "github.com/Iqblyh/recfood/user/domain/mocks"
	"github.com/Iqblyh/recfood/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userService userDomain.Service
	domain      userDomain.User
	userRepo    userMock.Repository
)

func TestMain(m *testing.M) {
	userService = service.NewUserService(&userRepo, &middlewares.ConfigJWT{})
	domain = userDomain.User{
		Id:       1,
		Username: "iqblyh",
		Email:    "abc@mail.com",
		Password: "rahasia",
		Fullname: "Iqbal Yoga",
	}
	os.Exit(m.Run())
}


func TestGetId(t *testing.T) {

}

func TestInsertData(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userRepo.On("Save", mock.AnythingOfType("domain.User")).Return(1, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(domain, nil).Once()

		res, err := userService.InsertData(domain)

		assert.NoError(t, err)
		assert.Equal(t, "iqblyh", res.Username)
	})

}

func TestLogin(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		userRepo.On("GetByUsernamePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepo.On("GenerateToken", mock.AnythingOfType("int")).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc", nil).Once()

		res, err := userService.Login("iqblyh", "rahasia")

		assert.NoError(t, err)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc", res)
	})
}