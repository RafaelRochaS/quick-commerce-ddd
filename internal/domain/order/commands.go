package order

import (
	"context"
)

type PlaceOrderCmd struct {
	Order *Order
	Ctx   context.Context
}

func NewPlaceOrderCmd(order *Order, ctx context.Context) *PlaceOrderCmd {
	return &PlaceOrderCmd{Order: order, Ctx: ctx}
}
