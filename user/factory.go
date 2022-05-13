package user

import (
	"gorm.io/gorm"

	userApi "github.com/Iqblyh/recfood/user/handler/api"
	userRepoMysql "github.com/Iqblyh/recfood/user/repository/mysql"
	userService "github.com/Iqblyh/recfood/user/service"
)

func NewUserFactory(db *gorm.DB) (userHandler userApi.UserHandler) {
	userRepo := userRepoMysql.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo)
	userHandler = userApi.NewUserHandler(userServ)
	return
}
