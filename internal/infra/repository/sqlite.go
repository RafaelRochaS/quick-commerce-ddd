package repository

import (
	"log"

	"github.com/RafaelRochaS/quick-commerce-ddd/internal/domain/order"
)

type SQLite struct {
	conn string
}

func NewSQL() order.Repository {
	return &SQLite{}
}

func (S SQLite) PlaceOrder(order *order.Order) error {
	log.Println("PlaceOrder called:", order)
	log.Println("Order placed")

	return nil
}
