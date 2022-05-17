package culinaryApi

import (
	"fmt"
	"net/http"
	"strconv"

	culinaryDomain "github.com/Iqblyh/recfood/culinary/domain"
	shopDomain "github.com/Iqblyh/recfood/shop/domain"
	"github.com/labstack/echo/v4"
)

type CulinaryHandler struct {
	serviceCulinary culinaryDomain.Service
	serviceShop     shopDomain.Service
}

func NewCulinaryHandler(culinary culinaryDomain.Service, shop shopDomain.Service) CulinaryHandler {
	return CulinaryHandler{
		serviceCulinary: culinary,
		serviceShop:     shop,
	}
}

func (ch CulinaryHandler) InsertData(ctx echo.Context) error {
	req := RequestJSON{}
	shop_id, _ := strconv.Atoi(ctx.Param("id"))

	_, err := ch.serviceShop.GetDataByID(shop_id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	respData, err := ch.serviceCulinary.InsertData(shop_id, toDomain(req))
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

func (ch CulinaryHandler) EditData(ctx echo.Context) error {
	req := RequestJSON{}
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	respData, err := ch.serviceCulinary.EditData(id, toDomain(req))
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

func (ch CulinaryHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := ch.serviceCulinary.DeleteData(id)
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

func (ch CulinaryHandler) GetCulinarys(ctx echo.Context) error {
	res, err := ch.serviceCulinary.GetDatas()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success get datas",
		"data":    res,
		"rescode": http.StatusOK,
	})
}

func (ch CulinaryHandler) GetByCat(ctx echo.Context) error {
	cat := ctx.QueryParam("cat")
	res, err := ch.serviceCulinary.GetDataByCat(cat)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusNotFound,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success get datas",
		"data":    res,
		"rescode": http.StatusOK,
	})
}

func (ch CulinaryHandler) GetCulinary(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := ch.serviceCulinary.GetDataByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success get datas",
		"data":    fromDomain(res),
		"rescode": http.StatusOK,
	})
}

func (ch CulinaryHandler) SearchCulinary(ctx echo.Context) error {
	query := ctx.QueryParam("q")
	res, err := ch.serviceCulinary.GetByName(query)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success get datas",
		"data":    res,
		"rescode": http.StatusOK,
	})
}

func (ch CulinaryHandler) ShopCulinarys(ctx echo.Context) error {
	shop_id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := ch.serviceShop.GetDataByID(shop_id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusNotFound,
		})
	}
	res, err := ch.serviceCulinary.GetDataByShopID(shop_id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success get datas",
		"data":    res,
		"rescode": http.StatusOK,
	})
}

func (ch CulinaryHandler) Recomendation(ctx echo.Context) error {
	weather_cat, err := ch.serviceCulinary.CheckTemp()
	fmt.Println(weather_cat)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	res, err := ch.serviceCulinary.GetRecomendation(weather_cat)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"mesage":  "success get datas",
		"data":    fromDomain(res),
		"rescode": http.StatusOK,
	})
}
