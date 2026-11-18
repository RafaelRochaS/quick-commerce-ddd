package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RafaelRochaS/quick-commerce-ddd/internal/domain/order"
)

func main() {
	log.Println("Testing order aggregate:")

	productCost, err := order.NewValue(12, 5, "BRL")

	if err != nil {
		log.Fatal(err)
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
}
