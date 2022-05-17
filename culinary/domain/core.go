package domain

type Culinary struct {
	Shop_Id       int
	Category      string
	Culinary_Name string
	Price         int
	Description   string
	Weather_cat   string
}

type Weather struct {
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
}
