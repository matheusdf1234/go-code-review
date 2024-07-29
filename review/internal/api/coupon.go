package api

import (
	. "coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply(c *gin.Context) {
	apiReq := ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a *API) Create(c *gin.Context) {
	apiReq := CouponCreateDTO{}

	if err := c.BindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdCoupon, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	apiResponse := CouponCreateResponseDTO{
		ID:             createdCoupon.ID,
		Discount:       createdCoupon.Discount,
		Code:           createdCoupon.Code,
		MinBasketValue: createdCoupon.MinBasketValue,
	}
	c.JSON(http.StatusOK, apiResponse)
}

func (a *API) Get(c *gin.Context) {
	//TODO: if we look for a non existant coupon, things are exploding, I need to send back the proper request when that happens
	apiReq := CouponGetDTO{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		return
	}
	apiResponse := make([]CouponCreateResponseDTO, 0, len(coupons))

	//as a future improvement this "translation" of the  entity object to the DTO can be handled by a separate class
	//maybe inside of a folder called "Infrastructure"
	//in .Net we can use the "mapper" library, maybe there is a similar thing for go.
	for i := 0; i < len(coupons); i++ {
		couponToAdd := CouponCreateResponseDTO{
			ID:             coupons[i].ID,
			Discount:       coupons[i].Discount,
			Code:           coupons[i].Code,
			MinBasketValue: coupons[i].MinBasketValue,
		}
		apiResponse = append(apiResponse, couponToAdd)
	}

	c.JSON(http.StatusOK, apiResponse)
}
