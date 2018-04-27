package api

import (
	"fmt"
	"log"

	"github.com/jmataya/goeventtalk/commerce"
)

var orderCount int

type CartService struct{}

func (cs CartService) CreateCart(customerID int) string {
	orderCount++
	orderRef := fmt.Sprintf("BR000%d", orderCount)

	activity := commerce.Activity{
		Action: commerce.CreateCart,
		Payload: commerce.CreateCartActivity{
			CustomerID: customerID,
			OrderRef:   orderRef,
		},
	}

	log.Printf("%v", activity)
	return orderRef
}

func (cs CartService) AddLineItems(orderRef string, lineItems []commerce.LineItem) {
	activity := commerce.Activity{
		Action: commerce.AddLineItems,
		Payload: commerce.LineItemActivity{
			OrderRef:  orderRef,
			LineItems: lineItems,
		},
	}

	log.Printf("%v", activity)
}

func (cs CartService) AddShippingAddress(orderRef string, address commerce.Address) {
	activity := commerce.Activity{
		Action: commerce.AddShippingAddress,
		Payload: commerce.ShippingAddressActivity{
			OrderRef: orderRef,
			Address:  address,
		},
	}

	log.Printf("%v", activity)
}

func (cs CartService) AddPayment(orderRef string, payment commerce.Payment) {
	activity := commerce.Activity{
		Action: commerce.AddPayment,
		Payload: commerce.PaymentActivity{
			OrderRef: orderRef,
			Payment:  payment,
		},
	}

	log.Printf("%v", activity)
}
