package entity

//it makes no sense that we were sending the Basket object itself on the "applycoupon" endpoint request
// Firstly, because we are exposing the internal "Basket" object to the outside world
//but more importantly, because it makes no sense to send a request that contains
//the "AppliedDiscount"  andn "ApplicationSuccessful" parameters
//when, at that point in the applicaiton lifecycle, we have yet to do said operation.

//the only thing that the "ApplyCoupon" needs to know about is the code of the coupon and the value of the basket, nothing more
type ApplicationRequest struct {
	Code        string `json: "code"`
	BasketValue int    `json: "basketValue"`
	// Basket entity.Basket
}
