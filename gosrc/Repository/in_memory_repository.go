package repository

import (
	"couponManger-Api/gosrc/models"
	"errors"
	"sync"
)

var mu sync.Mutex

// InMemoryCouponRepository is a simple in-memory storage for coupons
type InMemoryCouponRepository struct {
	coupons   map[int]models.Coupon // Store coupons as the interface type
	idCounter int
}

// NewInMemoryCouponRepository creates a new in-memory repository
func NewInMemoryCouponRepository() *InMemoryCouponRepository {
	return &InMemoryCouponRepository{
		coupons:   make(map[int]models.Coupon),
		idCounter: 0, // Initialize the in-memory map
	}
}

// Ensure InMemoryCouponRepository implements CouponRepository
var _ CouponRepository = (*InMemoryCouponRepository)(nil)

// CreateCoupon adds a new coupon to the repository
func (r *InMemoryCouponRepository) CreateCoupon(coupon models.Coupon) error {
	mu.Lock()
	defer mu.Unlock()
	r.idCounter++                      // Increment the ID counter
	coupon.SetID(r.idCounter)          // Set the ID of the coupon
	r.coupons[coupon.GetID()] = coupon // Store the coupon using its ID
	return nil
}

// GetCouponByID retrieves a coupon by its ID
func (r *InMemoryCouponRepository) GetCouponByID(id int) (models.Coupon, error) {
	coupon, exists := r.coupons[id]
	if !exists {
		return nil, errors.New("coupon not found")
	}
	return coupon, nil // Return the coupon
}

// GetAllCoupons returns all coupons in the repository
func (r *InMemoryCouponRepository) GetAllCoupons() []models.Coupon {
	var allCoupons []models.Coupon
	for _, coupon := range r.coupons {
		allCoupons = append(allCoupons, coupon)
	}
	return allCoupons
}

// DeleteCoupon removes a coupon from the repository
func (r *InMemoryCouponRepository) DeleteCoupon(id int) error {
	if _, exists := r.coupons[id]; exists {
		delete(r.coupons, id)
		return nil
	}
	return errors.New("coupon not found")
}

// UpdateCoupon updates a coupon in the repository
func (r *InMemoryCouponRepository) UpdateCoupon(coupon models.Coupon) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := r.coupons[coupon.GetID()]; !exists {
		return errors.New("coupon not found")
	}
	r.coupons[coupon.GetID()] = coupon // Update the coupon using its ID
	return nil
}
