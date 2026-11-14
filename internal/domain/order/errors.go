package order

import "fmt"

type InvalidMoneyAmountError struct {
	Amount uint8
	Cents  uint8
}

func (e *InvalidMoneyAmountError) Error() string {
	return fmt.Sprintf("invalid amount %d, cents %d", e.Amount, e.Cents)
}

type InvalidCurrencyError struct {
	Currency string
}

func (e *InvalidCurrencyError) Error() string {
	return fmt.Sprintf("invalid currency: %s", e.Currency)
}
