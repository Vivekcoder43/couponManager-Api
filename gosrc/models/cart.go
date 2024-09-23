package models

type Cart struct {
	Items []CartItem `json:"items"`
}

func (c *Cart) TotalPrice() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
