package factory

import (
	"gorm.io/gorm"

	"github.com/Iqblyh/recfood/middlewares"
	userApi "github.com/Iqblyh/recfood/user/handler/api"
	userRepoMysql "github.com/Iqblyh/recfood/user/repository/mysql"
	userService "github.com/Iqblyh/recfood/user/service"
	"github.com/labstack/echo/v4/middleware"

	shopApi "github.com/Iqblyh/recfood/shop/handler/api"
	shopRepoMysql "github.com/Iqblyh/recfood/shop/repository/mysql"
	shopService "github.com/Iqblyh/recfood/shop/service"

	culinaryApi "github.com/Iqblyh/recfood/culinary/handler/api"
	culinaryRepoMysql "github.com/Iqblyh/recfood/culinary/repository/mysql"
	culinaryService "github.com/Iqblyh/recfood/culinary/service"
)

func Factory(db *gorm.DB) (userApi.UserHandler, shopApi.ShopHandler, culinaryApi.CulinaryHandler, middleware.JWTConfig) {
	configJWT := middlewares.ConfigJWT{
		SecretJWT: "1324",
	}
	jwtMiddleware := configJWT.Init()

	userRepo := userRepoMysql.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo, &configJWT)
	userHandler := userApi.NewUserHandler(userServ)

	shopRepo := shopRepoMysql.NewShopRepository(db)
	shopServ := shopService.NewShopService(shopRepo)
	shopHandler := shopApi.NewShopHandler(shopServ)

	culinaryRepo := culinaryRepoMysql.NewCulinaryRepository(db)
	culinaryServ := culinaryService.NewCulinaryService(culinaryRepo)
	culinaryHandler := culinaryApi.NewCulinaryHandler(culinaryServ, shopServ)

	return userHandler, shopHandler, culinaryHandler, jwtMiddleware
}
