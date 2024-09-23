package repository

import (
	"couponManger-Api/gosrc/models"
)

// CouponRepository defines the methods that a repository should implement
type CouponRepository interface {
	CreateCoupon(coupon models.Coupon) error
	GetCouponByID(id int) (models.Coupon, error)
	GetAllCoupons() []models.Coupon
	DeleteCoupon(id int) error
	UpdateCoupon(coupon models.Coupon) error
}
