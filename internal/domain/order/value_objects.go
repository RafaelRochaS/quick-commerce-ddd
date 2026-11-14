package order

import "slices"

var ValidCurrencies = []string{"BRL", "USD", "EUR", "GBP", "JPY"}

type Money struct {
	Amount   uint8
	Cents    uint8
	Currency string
}

func NewMoney(amount, cents uint8, currency string) (*Money, error) {
	if amount < 1 || cents < 1 {
		return nil, &InvalidMoneyAmountError{Cents: cents, Amount: amount}
	}

	if !slices.Contains(ValidCurrencies, currency) {
		return nil, &InvalidCurrencyError{Currency: currency}
	}

	return &Money{
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
