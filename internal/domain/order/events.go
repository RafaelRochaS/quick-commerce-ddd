package order

import "fmt"

type Created struct {
	Order Order
}

func NewCreatedEvent(o Order) *Created {
	return &Created{Order: o}
}

func (c Created) String() string {
	return fmt.Sprintf("OrderId: %d\nItemId: %d", c.Order.Id, c.Order.Item.Id)
}
