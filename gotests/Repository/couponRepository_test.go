package repository_test

import (
	repository "gosrc/Repository"
	"gosrc/models"
	"testing"
)

func TestCreateAndGetCoupon(t *testing.T) {
	// Create the in-memory repository
	repo := repository.NewInMemoryCouponRepository()

	// Create a sample coupon
	coupon := &models.CartCoupon{
		BaseCoupon: models.BaseCoupon{ID: 1, Type: "cart-wise"},
		Threshold:  100,
		Discount:   10,
	}

	// Test creating the coupon
	err := repo.CreateCoupon(coupon)
	if err != nil {
		t.Errorf("Failed to create coupon: %v", err)
	}

	// Test retrieving the coupon
	storedCoupon, err := repo.GetCouponByID(1)
	if err != nil {
		t.Errorf("Failed to retrieve coupon: %v", err)
	}

	// Assert the stored coupon is correct
	if storedCoupon.GetID() != 1 {
		t.Errorf("Expected coupon ID 1, got %v", storedCoupon.GetID())
	}
}

func TestDeleteCoupon(t *testing.T) {
	// Create the in-memory repository
	repo := repository.NewInMemoryCouponRepository()

	// Create a sample coupon
	coupon := &models.CartCoupon{
		BaseCoupon: models.BaseCoupon{ID: 1, Type: "cart-wise"},
		Threshold:  100,
		Discount:   10,
	}

	// Create the coupon
	err := repo.CreateCoupon(coupon)
	if err != nil {
		t.Errorf("Failed to create coupon: %v", err)
	}

	// Test deleting the coupon
	err = repo.DeleteCoupon(1)
	if err != nil {
		t.Errorf("Failed to delete coupon: %v", err)
	}

	// Test that the coupon is no longer in the repository
	_, err = repo.GetCouponByID(1)
	if err == nil {
		t.Errorf("Expected error when retrieving deleted coupon, got none")
	}
}

func TestGetCoupon_NotFound(t *testing.T) {
	// Create the in-memory repository
	repo := repository.NewInMemoryCouponRepository()

	// Test retrieving a non-existent coupon
	_, err := repo.GetCouponByID(999)
	if err == nil {
		t.Errorf("Expected error when retrieving non-existent coupon, got none")
	}
}
