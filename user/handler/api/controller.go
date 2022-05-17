package userApi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Iqblyh/recfood/user/domain"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service domain.Service
}

func NewUserHandler(service domain.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (uh UserHandler) Register(ctx echo.Context) error {
	req := RequestJSON{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	res, err := uh.service.InsertData(toDomain(req))

	if err != nil {
		return errors.New("hadeehh")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":    fromDomain(res),
		"message": "Succes Insert Data",
	})
}

func (uh UserHandler) Login(ctx echo.Context) error {
	req := RequestJSON{}

	if err := ctx.Bind(&req); err != nil {
		return errors.New("not found")
	}

	token, err := uh.service.Login(req.Username, req.Password)

	fmt.Println("data token ", token)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "anda tidak valid",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success",
		"response": http.StatusOK,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
