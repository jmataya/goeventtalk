package storage

import "github.com/jmataya/goeventtalk/commerce"

type Cart struct {
	CustomerID int
	OrderRef   string
	LineItems  []commerce.LineItem
	Address    commerce.Address
	Payment    commerce.Payment
}
