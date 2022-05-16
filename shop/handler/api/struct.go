package shopApi

import "github.com/Iqblyh/recfood/shop/domain"

type RequestJSON struct {
	Shop_Name string `json:"shop_name"`
	Address   string `json:"address"`
	Open      string `json:"open"`
	Close     string `json:"close"`
}

func toDomain(req RequestJSON) domain.Shop {
	return domain.Shop{
		Shop_Name: req.Shop_Name,
		Address:   req.Address,
		Open:      req.Open,
		Close:     req.Close,
	}
}

type ResponseJSON struct {
	Shop_Name string `json:"shop_name"`
	Address   string `json:"address"`
	Open      string `json:"open"`
	Close     string `json:"close"`
}

func fromDomain(domain domain.Shop) ResponseJSON {
	return ResponseJSON{
		Shop_Name: domain.Shop_Name,
		Address:   domain.Address,
		Open:      domain.Open,
		Close:     domain.Close,
	}
}
