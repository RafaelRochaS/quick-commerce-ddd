package order

type Order struct {
	id   Id
	item *Item
}

func NewOrder(item *Item) *Order {
	return &Order{
		id:   123456,
		item: item,
	}
}

func (o *Order) AggregateID() Id {
	return o.id
}
