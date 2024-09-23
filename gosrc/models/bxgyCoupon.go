// models/bxgyCoupon.go
package models

type BxGyCoupon struct {
	BaseCoupon
	BuyProducts     []BuyProduct `json:"buy_products"`
	GetProducts     []GetProduct `json:"get_products"`
	RepetitionLimit int          `json:"repetition_limit"`
}

func (b *BxGyCoupon) GetID() int {
	return b.ID
}

func (b *BxGyCoupon) SetID(id int) {
	b.ID = id
}

type BuyProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type GetProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// Apply logic for BxGy coupon
func (b *BxGyCoupon) Apply(cart *Cart) (float64, error) {
	// Implement logic for applying BxGy coupon here
	return 0, nil
}
