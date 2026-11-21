package order

type Repository interface {
	PlaceOrder(order *Order) error
}
