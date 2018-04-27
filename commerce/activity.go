package commerce

const (
	CreateCart         = "create_cart"
	AddLineItems       = "add_line_items"
	AddShippingAddress = "add_shipping_address"
	AddPayment         = "add_payment"
)

type Activity struct {
	Action  string
	Payload interface{}
}

type CreateCartActivity struct {
	CustomerID int
	OrderRef   string
}

type LineItemActivity struct {
	OrderRef  string
	LineItems []LineItem
}

type ShippingAddressActivity struct {
	OrderRef string
	Address  Address
}

type PaymentActivity struct {
	OrderRef string
	Payment  Payment
}
