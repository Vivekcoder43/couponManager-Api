package handlers_test

import (
	"bytes"
	"encoding/json"
	repository "gosrc/Repository"
	"gosrc/handlers"
	"gosrc/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCouponHandler_CartWise(t *testing.T) {
	// Create the repository, service, and handler
	repo := repository.NewInMemoryCouponRepository()
	service := services.NewCouponService(repo)
	handler := handlers.NewCouponHandler(service)

	// Create a sample Cart-Wise Coupon JSON
	coupon := map[string]interface{}{
		"type": "cart-wise",
		"details": map[string]interface{}{
			"threshold": 100,
			"discount":  10,
		},
	}

	// Convert the coupon map to JSON
	couponJSON, err := json.Marshal(coupon)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a new HTTP request with the coupon JSON
	req, err := http.NewRequest("POST", "/coupons", bytes.NewBuffer(couponJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the CreateCouponHandler
	handlerFunc := http.HandlerFunc(handler.CreateCouponHandler)
	handlerFunc.ServeHTTP(rr, req)

	// Check the response code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	expected := `{"message":"Coupon created successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
