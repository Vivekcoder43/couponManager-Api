package models

type Coupon interface {
	GetID() int
	SetID(id int)
	Apply(cart *Cart) (float64, error) // Apply discount logic for each coupon type
}
