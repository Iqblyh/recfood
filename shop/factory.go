package shop

import (
	shopApi "github.com/Iqblyh/recfood/shop/handler/api"
	shopRepoMysql "github.com/Iqblyh/recfood/shop/repository/mysql"
	shopService "github.com/Iqblyh/recfood/shop/service"
	"gorm.io/gorm"
)

func NewShopFactory(db *gorm.DB) (shopHandler shopApi.ShopHandler) {
	shopRepo := shopRepoMysql.NewShopRepository(db)
	shopServ := shopService.NewShopService(shopRepo)
	shopHandler = shopApi.NewShopHandler(shopServ)
	return
}
