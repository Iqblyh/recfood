package culinaryApi

import "github.com/Iqblyh/recfood/culinary/domain"

type RequestJSON struct {
	Category      string `json:"category"`
	Culinary_Name string `json:"culinary_name"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	Weather_Cat   string `json:"weather_cat"`
}

func toDomain(req RequestJSON) domain.Culinary {
	return domain.Culinary{
		Category:      req.Category,
		Culinary_Name: req.Culinary_Name,
		Price:         req.Price,
		Description:   req.Description,
		Weather_cat:   req.Weather_Cat,
	}
}

type ResponseJSON struct {
	Shop_Id       int    `json:"shop_id"`
	Category      string `json:"category"`
	Culinary_Name string `json:"culinary_name"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
}

func fromDomain(domain domain.Culinary) ResponseJSON {
	return ResponseJSON{
		Shop_Id:       domain.Shop_Id,
		Category:      domain.Category,
		Culinary_Name: domain.Culinary_Name,
		Price:         domain.Price,
		Description:   domain.Description,
	}
}
