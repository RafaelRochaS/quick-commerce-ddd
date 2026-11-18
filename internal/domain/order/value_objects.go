package order

import "slices"

var ValidCurrencies = []string{"BRL", "USD", "EUR", "GBP", "JPY"}

type Value struct {
	Amount   uint8
	Cents    uint8
	Currency string
}

func NewValue(amount, cents uint8, currency string) (*Value, error) {
	if amount < 1 || cents < 1 {
		return nil, &InvalidMoneyAmountError{Cents: cents, Amount: amount}
	}

	if !slices.Contains(ValidCurrencies, currency) {
		return nil, &InvalidCurrencyError{Currency: currency}
	}

	return &Value{
		Amount:   amount,
		Cents:    cents,
		Currency: currency,
	}, nil
}

type ShippingAddress struct {
	Country  string
	State    string
	City     string
	Street   string
	Number   int8
	Postcode string
}

func NewShippingAddress(country, state, city, street, postcode string, number int8) *ShippingAddress {
	return &ShippingAddress{
		Country:  country,
		State:    state,
		City:     city,
		Street:   street,
		Number:   number,
		Postcode: postcode,
	}
}

type Status int8

const (
	Placed Status = iota
	PaymentPending
	PaymentApproved
	PaymentRejected
	Shipped
	Finished
	Cancelled
)

func (s Status) String() string {
	switch s {
	case Placed:
		return "ORDER_PLACED"
	case PaymentPending:
		return "ORDER_PAYMENT_PENDING"
	case PaymentApproved:
		return "ORDER_PAYMENT_APPROVED"
	case PaymentRejected:
		return "ORDER_PAYMENT_REJECTED"
	case Shipped:
		return "ORDER_SHIPPED"
	case Finished:
		return "ORDER_FINISHED"
	case Cancelled:
		return "ORDER_CANCELLED"
	default:
		return "UNKNOWN_ORDER_STATUS"
	}
}

type Product struct {
	Name        string
	Description string
	Value       *Value
	SKU         string
}

func NewProduct(name, description string, value *Value, sku string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Value:       value,
		SKU:         sku,
	}
}
