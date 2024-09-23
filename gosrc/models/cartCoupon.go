package models

// CartCoupon applies a discount if the cart exceeds a threshold
type CartCoupon struct {
	BaseCoupon
	Threshold int     `json:"threshold"`
	Discount  float64 `json:"discount"`
}

func (c *CartCoupon) GetID() int {
	return c.ID
}

func (c *CartCoupon) SetID(id int) {
	c.ID = id

}

// Apply logic for cart-wise coupon
func (c *CartCoupon) Apply(cart *Cart) (float64, error) {
	totalCartPrice := cart.TotalPrice()
	if totalCartPrice >= float64(c.Threshold) {
		discount := totalCartPrice * (c.Discount / 100)
		return discount, nil
	}
	return 0, nil
}
