package order

type Order struct {
	Id   Id
	Item *Item
}

func NewOrder(item *Item) *Order {
	return &Order{
		Id:   123456,
		Item: item,
	}
}
