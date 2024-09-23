package services

import (
	repository "couponManager-Api/gosrc/Repository"
	"couponManager-Api/gosrc/models"
	"couponManager-Api/gosrc/services"
	"testing"
)

func TestCreateCartCoupon(t *testing.T) {
	// Create the in-memory repository and the service
	repo := repository.NewInMemoryCouponRepository()
	service := services.NewCouponService(repo)

	// Create a sample cart-wise coupon
	coupon := &models.CartCoupon{
		BaseCoupon: models.BaseCoupon{ID: 1, Type: "cart-wise"},
		Threshold:  100,
		Discount:   10,
	}

	// Test creating the coupon
	err := service.CreateCoupon(coupon)
	if err != nil {
		t.Errorf("Failed to create coupon: %v", err)
	}

	// Retrieve the coupon from the repository and check its properties
	storedCoupon, err := repo.GetCouponByID(1)
	if err != nil {
		t.Errorf("Failed to retrieve coupon: %v", err)
	}

	// Assert the retrieved coupon is the same
	if storedCoupon.GetID() != 1 {
		t.Errorf("Expected coupon ID 1, got %v", storedCoupon.GetID())
	}
}

func TestCreateProductCoupon(t *testing.T) {
	// Create the in-memory repository and the service
	repo := repository.NewInMemoryCouponRepository()
	service := services.NewCouponService(repo)

	// Create a sample product-wise coupon
	coupon := &models.ProductCoupon{
		BaseCoupon: models.BaseCoupon{ID: 2, Type: "product-wise"},
		ProductID:  101,
		Discount:   20,
	}

	// Test creating the coupon
	err := service.CreateCoupon(coupon)
	if err != nil {
		t.Errorf("Failed to create coupon: %v", err)
	}

	// Retrieve the coupon from the repository and check its properties
	storedCoupon, err := repo.GetCouponByID(2)
	if err != nil {
		t.Errorf("Failed to retrieve coupon: %v", err)
	}

	// Assert the retrieved coupon is the same
	if storedCoupon.GetID() != 2 {
		t.Errorf("Expected coupon ID 2, got %v", storedCoupon.GetID())
	}
}

func TestCreateBxGyCoupon(t *testing.T) {
	// Create the in-memory repository and the service
	repo := repository.NewInMemoryCouponRepository()
	service := services.NewCouponService(repo)

	// Create a sample BxGy coupon
	coupon := &models.BxGyCoupon{
		BaseCoupon: models.BaseCoupon{ID: 3, Type: "bxgy"},
		BuyProducts: []models.BuyProduct{
			{ProductID: 1, Quantity: 3},
			{ProductID: 2, Quantity: 2},
		},
		GetProducts: []models.GetProduct{
			{ProductID: 3, Quantity: 1},
		},
		RepetitionLimit: 2,
	}

	// Test creating the coupon
	err := service.CreateCoupon(coupon)
	if err != nil {
		t.Errorf("Failed to create coupon: %v", err)
	}

	// Retrieve the coupon from the repository and check its properties
	storedCoupon, err := repo.GetCouponByID(3)
	if err != nil {
		t.Errorf("Failed to retrieve coupon: %v", err)
	}

	// Assert the retrieved coupon is the same
	if storedCoupon.GetID() != 3 {
		t.Errorf("Expected coupon ID 3, got %v", storedCoupon.GetID())
	}
}
