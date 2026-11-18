package order

type Order struct {
	Item *Item
}

func NewOrder(item *Item) *Order {
	return &Order{
		Item: item,
	}
}
