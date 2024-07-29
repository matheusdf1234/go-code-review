package entity

type CouponCreateResponseDTO struct {
	ID             string `json:"id"`
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"minBasketValue"`
}
