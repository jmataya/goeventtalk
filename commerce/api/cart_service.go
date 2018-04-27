package api

import (
	"fmt"

	"github.com/jmataya/goeventtalk/commerce/events"

	"github.com/jmataya/goeventtalk/commerce"
)

var orderCount int

// CartService is a collection of functions designed to build a cart and
// create an order.
type CartService struct {
	producer  *events.Producer
	topic     string
	partition int32
}

// CreateCart generates a new cart for the customer and returns a reference to
// the Order ID.
func (cs CartService) CreateCart(customerID int) (string, error) {
	orderCount++
	orderRef := fmt.Sprintf("BR000%d", orderCount)

	activity := commerce.Activity{
		Action: commerce.CreateCart,
		Payload: commerce.CreateCartActivity{
			CustomerID: customerID,
			OrderRef:   orderRef,
		},
	}

	// log.Printf("%v", activity)
	// return orderRef, nil
	err := cs.producer.Produce(cs.topic, cs.partition, &activity)
	return orderRef, err
}

// AddLineItems adds one or more items to the cart.
func (cs CartService) AddLineItems(orderRef string, lineItems []commerce.LineItem) error {
	activity := commerce.Activity{
		Action: commerce.AddLineItems,
		Payload: commerce.LineItemActivity{
			OrderRef:  orderRef,
			LineItems: lineItems,
		},
	}

	// log.Printf("%v", activity)
	// return nil
	return cs.producer.Produce(cs.topic, cs.partition, &activity)
}

// AddShippingAddress adds or replaces the cart's shipping address.
func (cs CartService) AddShippingAddress(orderRef string, address commerce.Address) error {
	activity := commerce.Activity{
		Action: commerce.AddShippingAddress,
		Payload: commerce.ShippingAddressActivity{
			OrderRef: orderRef,
			Address:  address,
		},
	}

	// log.Printf("%v", activity)
	// return nil
	return cs.producer.Produce(cs.topic, cs.partition, &activity)
}

// AddPayment adds or replaces the cart't payment method.
func (cs CartService) AddPayment(orderRef string, payment commerce.Payment) error {
	activity := commerce.Activity{
		Action: commerce.AddPayment,
		Payload: commerce.PaymentActivity{
			OrderRef: orderRef,
			Payment:  payment,
		},
	}

	// log.Printf("%v", activity)
	// return nil
	return cs.producer.Produce(cs.topic, cs.partition, &activity)
}
