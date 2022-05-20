package main

import (
	"github.com/Iqblyh/recfood/config"
	"github.com/Iqblyh/recfood/routes"
)

func main() {
	config.Init()
	db := config.DBInit()
	e := routes.New(db)

	e.Start(":8080")
}
