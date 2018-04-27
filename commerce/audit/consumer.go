package audit

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmataya/goeventtalk/commerce"
	"github.com/jmataya/goeventtalk/commerce/events"
)

var customers = map[int]string{
	1: "Jeff Mataya",
	2: "Adil Wali",
	3: "Bree Swineford",
}

var orders = map[string]int{}

func Run() error {
	consumer, err := events.NewConsumer("localhost:29092", "cartsAudit", "earliest")
	if err != nil {
		return err
	}

	fmt.Println("-----------------")
	fmt.Println("Audit Log Started")
	fmt.Println("-----------------")
	fmt.Println("")

	return consumer.Consume("carts", kafka.PartitionAny, handleActivity)
}

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

	orders[cartActivity.OrderRef] = cartActivity.CustomerID
	customerName := customers[cartActivity.CustomerID]

	fmt.Printf("%s created a new cart with reference %s\n", customerName, cartActivity.OrderRef)
	return nil
}

func handleAddLineItems(a *commerce.Activity) error {
	lineItemActivity := new(commerce.LineItemActivity)
	if err := getPayload(a, lineItemActivity); err != nil {
		return err
	}

	customerID := orders[lineItemActivity.OrderRef]
	customerName := customers[customerID]

	for _, lineItem := range lineItemActivity.LineItems {
		fmt.Printf(
			"%s %d unit of %s to cart %s\n",
			customerName,
			lineItem.Quantity,
			lineItem.SKU,
			lineItemActivity.OrderRef,
		)
	}

	return nil
}

func handleAddShippingAddress(a *commerce.Activity) error {
	shippingActivity := new(commerce.ShippingAddressActivity)
	if err := getPayload(a, shippingActivity); err != nil {
		return err
	}

	customerID := orders[shippingActivity.OrderRef]
	customerName := customers[customerID]

	fmt.Printf(
		"%s added the address %s, %s %s, %s %s to the cart %s\n",
		customerName,
		shippingActivity.Address.Street1,
		shippingActivity.Address.Street2,
		shippingActivity.Address.City,
		shippingActivity.Address.State,
		shippingActivity.Address.PostalCode,
		shippingActivity.OrderRef,
	)

	return nil
}

func handleAddPayment(a *commerce.Activity) error {
	paymentActivity := new(commerce.PaymentActivity)
	if err := getPayload(a, paymentActivity); err != nil {
		return err
	}

	customerID := orders[paymentActivity.OrderRef]
	customerName := customers[customerID]

	fmt.Printf(
		"%s added payment method with expiration %d/%d to cart %s\n",
		customerName,
		paymentActivity.Payment.ExpMonth,
		paymentActivity.Payment.ExpYear,
		paymentActivity.OrderRef,
	)

	return nil
}

func getPayload(a *commerce.Activity, out interface{}) error {
	if err := json.Unmarshal(a.Payload.([]byte), out); err != nil {
		return err
	}

	return nil
}
