package middlewares

import (
	"time"

	"github.com/Iqblyh/recfood/config"
	errorConv "github.com/Iqblyh/recfood/helper/error"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigJWT struct {
	SecretJWT string
}

func (cj ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte(cj.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(err error, c echo.Context) error {
			return errorConv.Conversion(err)
		}),
	}
}

//token
func (cj ConfigJWT) CreateToken(user_id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claimToken.SignedString([]byte(config.Conf.JWTSecret))

	return token, err
}
