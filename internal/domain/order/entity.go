package order

type Id int64

type Item struct {
	Id      Id
	Product *Product
	Status  Status
	Address *ShippingAddress
}

func NewItem(id Id, product *Product, address *ShippingAddress) *Item {
	return &Item{Id: id,
		Product: product,
		Address: address}
}
