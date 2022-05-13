package config

import (
	"fmt"
	"os"

	repository "github.com/Iqblyh/recfood/user/repository/mysql"
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
	os.Setenv("DBNAME", "recfood")
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "1324")
	os.Setenv("DBHOST", "127.0.0.1")
	os.Setenv("DBPORT", "3306")
	os.Setenv("JWTSecret", "1324")
	Conf = Config{
		DBNAME:    os.Getenv("DBNAME"),
		DBUSER:    os.Getenv("DBUSER"),
		DBPASS:    os.Getenv("DBPASS"),
		DBHOST:    os.Getenv("DBHOST"),
		DBPORT:    os.Getenv("DBPORT"),
		JWTSecret: os.Getenv("JWTSECRET"),
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
	DB.AutoMigrate(&repository.User{})

	return
}
