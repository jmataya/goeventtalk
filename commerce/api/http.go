package api

import (
	"net/http"

	"github.com/jmataya/goeventtalk/commerce"
	"github.com/labstack/echo"
)

func Run() {
	// CartService instance
	cs := new(CartService)

	// Echo instance
	e := echo.New()

	// Routes
	e.POST("/cart", createCart(cs))
	e.PATCH("/cart/:orderRef/line-items", addLineItems(cs))
	e.PATCH("/cart/:orderRef/shipping-address", addShippingAddress(cs))
	e.PATCH("/cart/:orderRef/payment-method", addPayment(cs))

	// Start server
	e.Logger.Fatal(e.Start(":21337"))
}

type errMessage struct {
	Error error
}

func createCart(cs *CartService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload struct {
			CustomerID int
		}

		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, errMessage{err})
		}

		orderRef := cs.CreateCart(payload.CustomerID)
		return c.JSON(http.StatusCreated, struct {
			OrderRef string
		}{orderRef})
	}
}

func addLineItems(cs *CartService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload struct {
			LineItems []commerce.LineItem
		}

		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, errMessage{err})
		}

		orderRef := c.Param("orderRef")
		cs.AddLineItems(orderRef, payload.LineItems)

		return c.JSON(http.StatusOK, payload)
	}
}

func addShippingAddress(cs *CartService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload struct {
			Address commerce.Address
		}

		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, errMessage{err})
		}

		orderRef := c.Param("orderRef")
		cs.AddShippingAddress(orderRef, payload.Address)

		return c.JSON(http.StatusOK, payload)
	}
}

func addPayment(cs *CartService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload struct {
			Payment commerce.Payment
		}

		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, errMessage{err})
		}

		orderRef := c.Param("orderRef")
		cs.AddPayment(orderRef, payload.Payment)

		return c.JSON(http.StatusOK, payload)
	}
}
