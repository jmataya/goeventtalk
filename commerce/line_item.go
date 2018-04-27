package commerce

// LineItem is an item purchased in an order.
type LineItem struct {
	SKU      string  `json:"sku"`
	Quantity int     `json:"qty"`
	Price    float64 `json:"price"`
}
