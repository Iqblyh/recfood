package middlewares

import (
	errorConv "github.com/Iqblyh/recfood/helper/error"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigJWT struct {
	SecretJWT string
}

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func (cj ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(err error, c echo.Context) error {
			return errorConv.Conversion(err)
		}),
	}
}

//token
func (jwtConf *ConfigJWT) GenerateToken(userID int) (string, error) {
	claims := JwtCustomClaims{
		Id: userID,
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
