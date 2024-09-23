package services

import (
	models2 "couponManger-Api/gosrc/models"
	"couponManger-Api/gosrc/repository"
)

// CouponService handles the business logic for coupons
type CouponService struct {
	Repo repository.CouponRepository // Use the CouponRepository interface from the repository package
}

// NewCouponService creates a new CouponService
func NewCouponService(repo repository.CouponRepository) *CouponService {
	return &CouponService{Repo: repo}
}

// CreateCoupon creates a new coupon
func (s *CouponService) CreateCoupon(coupon models2.Coupon) error {
	return s.Repo.CreateCoupon(coupon)
}

// GetCouponByID retrieves a coupon by its ID
func (s *CouponService) GetCouponByID(id int) (models2.Coupon, error) {
	return s.Repo.GetCouponByID(id)
}

// GetAllCoupons retrieves all coupons
func (s *CouponService) GetAllCoupons() []models2.Coupon {
	return s.Repo.GetAllCoupons()
}

// DeleteCoupon deletes a coupon by its ID
func (s *CouponService) DeleteCoupon(id int) error {
	return s.Repo.DeleteCoupon(id)
}

// ApplyCoupon applies a coupon to the cart by ID
func (s *CouponService) ApplyCoupon(id int, cart *models2.Cart) (float64, error) {
	coupon, err := s.Repo.GetCouponByID(id)
	if err != nil {
		return 0, err
	}

	// Apply the coupon logic to the cart
	discount, err := coupon.Apply(cart)
	if err != nil {
		return 0, err
	}
	return discount, nil
}

func (s *CouponService) UpdateCoupon(coupon models2.Coupon) error {
	return s.Repo.UpdateCoupon(coupon)
}
