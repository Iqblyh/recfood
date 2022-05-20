package routes

import (
	"github.com/Iqblyh/recfood/shop"
	"github.com/Iqblyh/recfood/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	//user
	user := user.NewUserFactory(db)
	e.POST("/user", user.InsertData)

	//shop
	shop := shop.NewShopFactory(db)
	e.POST("/shop", shop.InsertData)
	e.GET("/shop/:id", shop.GetShop)
	e.GET("/shop", shop.GetShops)
	e.PUT("/shop/:id", shop.EditData)
	e.DELETE("/shop/:id", shop.DeleteData)
	return e
}
