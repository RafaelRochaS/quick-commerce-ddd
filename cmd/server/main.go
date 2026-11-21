package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RafaelRochaS/quick-commerce-ddd/internal/domain/order"
	"github.com/RafaelRochaS/quick-commerce-ddd/internal/infra/repository"
	"github.com/RafaelRochaS/quick-commerce-ddd/internal/infra/services"
)

func main() {
	log.Println("Testing order aggregate:")

	createdOrder, err := createOrder()

	if err != nil {
		log.Fatal("failed to create order:", err)
	}

	err = placeOrder(createdOrder)

	if err != nil {
		log.Fatal("failed to handle order:", err)
	}

}

func createOrder() (*order.Order, error) {
	productCost, err := order.NewValue(12, 5, "BRL")

	if err != nil {
		return nil, err
	}

	orderAddress := order.NewShippingAddress("Brasil",
		"SP",
		"Cidade",
		"Rua",
		"123123-101",
		1)
	productOrder := order.NewProduct("ProdutoA",
		"produto de teste",
		productCost,
		"123-123-123-123")
	itemOrder := order.NewItem(123, productOrder, orderAddress)
	newOrder := order.NewOrder(itemOrder)

	b, _ := json.MarshalIndent(newOrder, "", "  ")
	fmt.Println()

	log.Printf("Created new order: %s", string(b))

	return newOrder, nil
}

func placeOrder(readyOrder *order.Order) error {
	handler := order.NewPlaceOrderHandler(repository.NewSQL(), services.NewObsService())
	cmd := order.NewPlaceOrderCmd(readyOrder, context.Background())

	return handler.Handle(cmd)
}
