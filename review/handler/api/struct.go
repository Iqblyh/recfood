package reviewApi

import "github.com/Iqblyh/recfood/review/domain"

type RequestJSON struct {
	Rating int    `json:"rating"`
	Review string `json:"review"`
}

func toDomain(req RequestJSON) domain.Review {
	return domain.Review{
		Rating: req.Rating,
		Review: req.Review,
	}
}

type ResponseJSON struct {
	Culinary_Id int    `json:"culinary_id"`
	User_Id     int    `json:"user_id"`
	Rating      int    `json:"rating"`
	Review      string `json:"review"`
}

func fromDomain(domain domain.Review) ResponseJSON {
	return ResponseJSON{
		Culinary_Id: domain.Culinary_Id,
		User_Id:     domain.User_Id,
		Rating:      domain.Rating,
		Review:      domain.Review,
	}
}
