package models

// BaseCoupon represents the base structure for all coupons
type BaseCoupon struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// GetID returns the ID of the coupon
func (b *BaseCoupon) GetID() int {
	return b.ID
}

// SetID sets the ID of the coupon
func (b *BaseCoupon) SetID(id int) {
	b.ID = id
}
