package config

import (
	"fmt"

	repositoryCulinary "github.com/Iqblyh/recfood/culinary/repository/mysql"
	repositoryReview "github.com/Iqblyh/recfood/review/repository/mysql"
	repositoryShop "github.com/Iqblyh/recfood/shop/repository/mysql"
	repositoryUser "github.com/Iqblyh/recfood/user/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string

	JWTSecret string
}

var Conf Config

func Init() {
	Conf = Config{
		DBNAME:    "recfood",
		DBUSER:    "root",
		DBPASS:    "",
		DBHOST:    "127.0.0.1",
		DBPORT:    "3306",
		JWTSecret: "1324",
	}
	fmt.Printf("%+v", Conf)
}

func DBInit() (DB *gorm.DB) {
	DB, _ = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				Conf.DBUSER,
				Conf.DBPASS,
				Conf.DBHOST,
				Conf.DBPORT,
				Conf.DBNAME,
			),
		),
	)
	DB.AutoMigrate(&repositoryUser.User{}, &repositoryShop.Shop{}, &repositoryCulinary.Culinary{}, &repositoryReview.Review{})

	return
}
