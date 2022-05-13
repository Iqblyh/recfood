package userApi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Iqblyh/recfood/config"
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

func (uh UserHandler) HealthCheck(ctx echo.Context) error {
	fmt.Println(config.Conf)
	return ctx.HTML(http.StatusOK, "OK HealtCheck + ")
}

func (uh UserHandler) InsertData(ctx echo.Context) error {
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

