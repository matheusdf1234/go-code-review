package entity

type CouponCreateDTO struct {
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"minbasketvalue"`
}
