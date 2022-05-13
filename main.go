package main

import (
	"github.com/Iqblyh/recfood/config"
	"github.com/Iqblyh/recfood/user"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Init()
	db := config.DBInit()

	user := user.NewUserFactory(db)

	//routing
	e := echo.New()

	e.GET("/healthceck", user.HealthCheck)
	e.POST("/user", user.InsertData)

	e.Start(":8080")
}
