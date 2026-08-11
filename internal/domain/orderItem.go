package domain

type OrderItem struct {
	id        int64
	productID int64
	name      string
	price     float64
	quantity  int
}

func (oi OrderItem) ID() int64 {
	return oi.id
}

func (oi OrderItem) ProductID() int64 {
	return oi.productID
}

func (oi OrderItem) Name() string {
	return oi.name
}

func (oi OrderItem) Price() float64 {
	return oi.price
}

func (oi OrderItem) Quantity() int {
	return oi.quantity
}

func NewOrderItem(id int64, product Product, quantity int) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}
	return &OrderItem{
		id:        id,
		productID: product.ID,
		name:      product.Name,
		price:     product.Price,
		quantity:  quantity,
	}, nil
}
