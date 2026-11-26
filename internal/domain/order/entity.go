package order

type Id int64

type Item struct {
	id      Id
	product *Product
	status  Status
	address *ShippingAddress
}

func NewItem(id Id, product *Product, address *ShippingAddress) *Item {
	return &Item{id: id,
		product: product,
		address: address}
}
