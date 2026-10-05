package task14

type Item struct {
	name     string
	price    float64
	quantity int
}

func NewItem(name string, price float64, quantity int) Item {
	return Item{name: name, price: price, quantity: quantity}
}

func (i Item) Name() string {
	return i.name
}

func (i Item) Price() float64 {
	return i.price
}

func (i Item) Quantity() int {
	return i.quantity
}
