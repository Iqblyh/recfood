package reviewApi

import (
	"fmt"
	"net/http"
	"strconv"

	culinaryDomain "github.com/Iqblyh/recfood/culinary/domain"
	"github.com/Iqblyh/recfood/middlewares"
	reviewDomain "github.com/Iqblyh/recfood/review/domain"
	userDomain "github.com/Iqblyh/recfood/user/domain"
	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	serviceReview   reviewDomain.Service
	serviceCulinary culinaryDomain.Service
	serviceUser     userDomain.Service
}

func NewReviewHandler(review reviewDomain.Service, culinary culinaryDomain.Service, user userDomain.Service) ReviewHandler {
	return ReviewHandler{
		serviceReview:   review,
		serviceCulinary: culinary,
		serviceUser:     user,
	}
}

func (rh ReviewHandler) InsertData(ctx echo.Context) error {
	fmt.Println("wkwkwkwk")
	req := RequestJSON{}
	claims := middlewares.GetUser(ctx)
	culinary_id, _ := strconv.Atoi(ctx.Param("id"))

	_, err := rh.serviceCulinary.GetDataByID(culinary_id)

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
	// userId, _ := strconv.Atoi(user_id)
	fmt.Println("claimsid : ", claims)
	respData, err := rh.serviceReview.InsertData(culinary_id, claims.Id, toDomain(req))
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

func (rh ReviewHandler) EditData(ctx echo.Context) error {
	req := RequestJSON{}
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	respData, err := rh.serviceReview.EditData(id, toDomain(req))
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

func (rh ReviewHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := rh.serviceReview.DeleteData(id)
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

func (rh ReviewHandler) GetReviews(ctx echo.Context) error {
	culinary_id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := rh.serviceReview.GetDatas(culinary_id)
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

func (rh ReviewHandler) GetUserReviews(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := rh.serviceReview.GetUserDatas(id)
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

func (rh ReviewHandler) GetReview(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res, err := rh.serviceReview.GetDataByID(id)
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
