package service

import (
	"coupon_service/internal/service/entity"
	. "coupon_service/internal/service/entity"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	FindByCode(string) (*Coupon, error)
	Save(Coupon) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ApplyCoupon(basketValue int, code string) (b *Basket, e error) {
	b = new(entity.Basket)
	b.Value = basketValue
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
		return b, nil
	}
	if b.Value == 0 {
		return
	}

	return nil, fmt.Errorf("tried to apply discount to negative value")
}

func (s Service) CreateCoupon(discount int, code string, minBasketValue int) (*Coupon, error) {
	coupon := Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (s Service) GetCoupons(codes []string) ([]Coupon, error) {
	coupons := make([]Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("\n The coupon with code: %s and index: %d could not be found ", code, idx)
			} else {
				e = fmt.Errorf("\n %w; The coupon with code: %s and index: %d could not be found ", e, code, idx)
			}

		} else {

			coupons = append(coupons, *coupon)
		}

	}

	return coupons, e
}
