package model

type Item struct {
	ID           int     `json:"ID"`
	ItemCode     string  `json:"Item Code"`
	ItemName     string  `json:"Item Name"`
	BuyingPrice  float64 `json:"Buying Price"`
	SellingPrice float64 `json:"Selling Price"`
	ItemAmount   int     `json:"Item Amount"`
	Pieces       string  `json:"Pieces"`
}
