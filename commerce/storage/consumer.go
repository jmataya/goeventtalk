package storage

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmataya/goeventtalk/commerce"
	"github.com/jmataya/goeventtalk/commerce/events"
)

var carts map[string]Cart

func handleActivity(a *commerce.Activity) error {
	switch a.Action {
	case commerce.CreateCart:
		return handleCreateCart(a)
	case commerce.AddLineItems:
		return handleAddLineItems(a)
	case commerce.AddShippingAddress:
		return handleAddShippingAddress(a)
	case commerce.AddPayment:
		return handleAddPayment(a)
	}

	return nil
}

func handleCreateCart(a *commerce.Activity) error {
	cartActivity := new(commerce.CreateCartActivity)
	if err := getPayload(a, cartActivity); err != nil {
		return err
	}

	cart := Cart{
		OrderRef:   cartActivity.OrderRef,
		CustomerID: cartActivity.CustomerID,
	}

	carts[cartActivity.OrderRef] = cart

	return logCart("Cart %s created!", cart)
}

func handleAddLineItems(a *commerce.Activity) error {
	lineItemActivity := new(commerce.LineItemActivity)
	if err := getPayload(a, lineItemActivity); err != nil {
		return err
	}

	cart := carts[lineItemActivity.OrderRef]
	cart.LineItems = lineItemActivity.LineItems
	carts[lineItemActivity.OrderRef] = cart

	return logCart("Line items added to cart %s!", cart)
}

func handleAddShippingAddress(a *commerce.Activity) error {
	shippingActivity := new(commerce.ShippingAddressActivity)
	if err := getPayload(a, shippingActivity); err != nil {
		return err
	}

	cart := carts[shippingActivity.OrderRef]
	cart.Address = shippingActivity.Address
	carts[shippingActivity.OrderRef] = cart

	return logCart("Shipping address added to cart %s!", cart)
}

func handleAddPayment(a *commerce.Activity) error {
	paymentActivity := new(commerce.PaymentActivity)
	if err := getPayload(a, paymentActivity); err != nil {
		return err
	}

	cart := carts[paymentActivity.OrderRef]
	cart.Payment = paymentActivity.Payment
	carts[paymentActivity.OrderRef] = cart

	return logCart("Order %s completed!", cart)
}

func getPayload(a *commerce.Activity, out interface{}) error {
	if err := json.Unmarshal(a.Payload.([]byte), out); err != nil {
		return err
	}

	return nil
}

func logCart(message string, c Cart) error {
	msg := fmt.Sprintf(message, c.OrderRef)
	fmt.Printf("%s\n", msg)
	fmt.Printf("%+v\n", c)
	return nil
}

func Run() error {
	carts = map[string]Cart{}

	consumer, err := events.NewConsumer("localhost:29092", "cartsStorage", "earliest")
	if err != nil {
		return err
	}

	return consumer.Consume("carts", kafka.PartitionAny, handleActivity)
}
