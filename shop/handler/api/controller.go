package shopApi

import (
	"github.com/Iqblyh/recfood/shop/domain"
)

type ShopHandler struct {
	service domain.Service
}

func NewShopHandler(service domain.Service) ShopHandler {
	return ShopHandler{
		service: service,
	}
}
