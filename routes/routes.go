package routes

import (
	"github.com/Iqblyh/recfood/factory"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	user, shop, culinary, jwtMiddleware := factory.Factory(db)

	//user
	e.POST("/register", user.Register)
	e.POST("/login", user.Login)

	//shop
	shopGroup := e.Group("shop")
	shopGroup.GET("/:id", shop.GetShop)
	shopGroup.GET("/all", shop.GetShops)

	authShop := e.Group("shop")
	authShop.Use(middleware.JWTWithConfig(jwtMiddleware))
	authShop.POST("/create", shop.InsertData)
	authShop.PUT("/edit/:id", shop.EditData)
	authShop.DELETE("/delete/:id", shop.DeleteData)

	//culinary
	culinaryGroup := e.Group("culinary")
	culinaryGroup.GET("/recomendation", culinary.Recomendation)
	culinaryGroup.GET("/shop/:id", culinary.ShopCulinarys)
	culinaryGroup.GET("/:id", culinary.GetCulinary)
	culinaryGroup.GET("/all", culinary.GetCulinarys)
	culinaryGroup.GET("/", culinary.GetByCat)
	culinaryGroup.GET("/search", culinary.SearchCulinary)

	authCulinary := e.Group("culinary")
	authCulinary.Use(middleware.JWTWithConfig(jwtMiddleware))
	authCulinary.POST("/shop/:id/create", culinary.InsertData)
	authCulinary.PUT("/edit/:id", culinary.EditData)
	authCulinary.DELETE("/delete/:id", culinary.DeleteData)
	return e
}
