package handlers

import (
	models2 "couponManger-Api/gosrc/models"
	"couponManger-Api/gosrc/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CouponHandler handles HTTP requests for coupons
type CouponHandler struct {
	Service *services.CouponService
}

// NewCouponHandler creates a new CouponHandler
func NewCouponHandler(service *services.CouponService) *CouponHandler {
	return &CouponHandler{Service: service}
}

// CreateCouponHandler handles the creation of a new coupon
func (h *CouponHandler) CreateCouponHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct for the base payload
	var payload struct {
		Type    string          `json:"type"`
		Details json.RawMessage `json:"details"`
	}

	// Decode the base payload (type and details)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid coupon data", http.StatusBadRequest)
		return
	}

	// Switch on the type of the coupon
	switch payload.Type {
	case "cart-wise":
		// Handle cart-wise coupon
		var details models2.CartCoupon
		if err := json.Unmarshal(payload.Details, &details); err != nil {
			http.Error(w, "Invalid cart-wise details", http.StatusBadRequest)
			return
		}
		// Set the type for the coupon
		details.BaseCoupon.Type = "cart-wise"
		err = h.Service.CreateCoupon(&details)

	case "product-wise":
		// Handle product-wise coupon
		var details models2.ProductCoupon
		if err := json.Unmarshal(payload.Details, &details); err != nil {
			http.Error(w, "Invalid product-wise details", http.StatusBadRequest)
			return
		}
		// Set the type for the coupon
		details.BaseCoupon.Type = "product-wise"
		err = h.Service.CreateCoupon(&details)

	case "bxgy":
		// Handle bxgy coupon
		var details models2.BxGyCoupon
		if err := json.Unmarshal(payload.Details, &details); err != nil {
			http.Error(w, "Invalid bxgy details", http.StatusBadRequest)
			return
		}
		// Set the type for the coupon
		details.BaseCoupon.Type = "bxgy"
		err = h.Service.CreateCoupon(&details)

	default:
		// Invalid coupon type
		http.Error(w, "Invalid coupon type", http.StatusBadRequest)
		return
	}

	// Handle any error during coupon creation
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Coupon created successfully"})
}

// GetCouponByIDHandler retrieves a coupon by ID
func (h *CouponHandler) GetCouponByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	coupon, err := h.Service.GetCouponByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(coupon)
}

// GetAllCouponsHandler retrieves all coupons
func (h *CouponHandler) GetAllCouponsHandler(w http.ResponseWriter, r *http.Request) {
	coupons := h.Service.GetAllCoupons()
	json.NewEncoder(w).Encode(coupons)
}

// DeleteCouponHandler deletes a coupon by ID
func (h *CouponHandler) DeleteCouponHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteCoupon(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ApplyCouponHandler applies a coupon to the cart
func (h *CouponHandler) ApplyCouponHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	var cart models2.Cart
	err = json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, "Invalid cart data", http.StatusBadRequest)
		return
	}

	discount, err := h.Service.ApplyCoupon(id, &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"discount": discount})
}

func (h *CouponHandler) UpdateCouponHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid coupon ID", http.StatusBadRequest)
		return
	}

	// Define a struct for the base payload
	var payload struct {
		Type    string          `json:"type"`
		Details json.RawMessage `json:"details"`
	}

	// Decode the base payload (type and details)
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid coupon data", http.StatusBadRequest)
		return
	}

	// Switch on the type of the coupon
	switch payload.Type {
	case "cart-wise":
		// Handle cart-wise coupon
		var details models2.CartCoupon
		if err := json.Unmarshal(payload.Details, &details); err != nil {
			http.Error(w, "Invalid cart-wise details", http.StatusBadRequest)
			return
		}
		details.SetID(id)
		err = h.Service.UpdateCoupon(&details)

	case "product-wise":
		// Handle product-wise coupon
		var details models2.ProductCoupon
		if err := json.Unmarshal(payload.Details, &details); err != nil {
			http.Error(w, "Invalid product-wise details", http.StatusBadRequest)
			return
		}
		details.SetID(id)
		err = h.Service.UpdateCoupon(&details)

	case "bxgy":
		// Handle bxgy coupon
		var details models2.BxGyCoupon
		if err := json.Unmarshal(payload.Details, &details); err != nil {
			http.Error(w, "Invalid bxgy details", http.StatusBadRequest)
			return
		}
		details.SetID(id)
		err = h.Service.UpdateCoupon(&details)

	default:
		// Invalid coupon type
		http.Error(w, "Invalid coupon type", http.StatusBadRequest)
		return
	}

	// Handle any error during coupon update
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Coupon updated successfully"})
}
