package shopApi

import (
	"net/http"
	"strconv"

	"github.com/Iqblyh/recfood/shop/domain"
	"github.com/labstack/echo/v4"
)

type ShopHandler struct {
	service domain.Service
}

func NewShopHandler(service domain.Service) ShopHandler {
	return ShopHandler{
		service: service,
	}
}

func (sh ShopHandler) InsertData(ctx echo.Context) error {
	req := RequestJSON{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	respData, err := sh.service.InsertData(toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success insert data",
		"rescode": http.StatusOK,
		"data":    fromDomain(respData),
	})
}

func (sh ShopHandler) EditData(ctx echo.Context) error {
	req := RequestJSON{}
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	respData, err := sh.service.EditData(id, toDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success update data",
		"rescode": http.StatusOK,
		"data":    fromDomain(respData),
	})
}

func (sh ShopHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := sh.service.DeleteData(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success delete data",
		"rescode": http.StatusOK,
	})
}
