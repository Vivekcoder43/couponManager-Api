package main

import (
	"couponManger-Api/gosrc/handlers"
	"couponManger-Api/gosrc/services"
	repository "couponManger-Api/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the in-memory repository
	couponRepo := repository.NewInMemoryCouponRepository()

	// Initialize the coupon service
	couponService := services.NewCouponService(couponRepo)

	// Initialize the coupon handler
	couponHandler := handlers.NewCouponHandler(couponService)

	// Set up the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/coupons", couponHandler.CreateCouponHandler).Methods("POST")
	r.HandleFunc("/coupons/{id}", couponHandler.GetCouponByIDHandler).Methods("GET")
	r.HandleFunc("/coupons", couponHandler.GetAllCouponsHandler).Methods("GET")
	r.HandleFunc("/coupons/{id}", couponHandler.DeleteCouponHandler).Methods("DELETE")
	r.HandleFunc("/apply-coupon/{id}", couponHandler.ApplyCouponHandler).Methods("POST")
	r.HandleFunc("/coupons/{id}", couponHandler.UpdateCouponHandler).Methods("PUT")
	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
