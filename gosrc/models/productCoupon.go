package models

// ProductCoupon applies a discount to a specific product
type ProductCoupon struct {
	BaseCoupon
	ProductID int     `json:"product_id"`
	Discount  float64 `json:"discount"`
}

func (p *ProductCoupon) GetID() int {
	return p.ID
}

func (p *ProductCoupon) SetID(id int) {
	p.ID = id
}
func (p *ProductCoupon) Apply(cart *Cart) (float64, error) {
	return p.ApplytoItem(&cart.Items[0])
}

// Apply logic for product-wise coupon to a single product
func (p *ProductCoupon) ApplytoItem(item *CartItem) (float64, error) {
	if item.ProductID == p.ProductID {
		totalPrice := item.Price * float64(item.Quantity)
		discount := totalPrice * (p.Discount / 100)
		return discount, nil
	}
	return 0, nil
}
