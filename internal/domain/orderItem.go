package domain

type OrderItem struct {
	ID        int64
	ProductID int64
	Name      string
	Price     float64
	Quantity  int
}

func NewOrderItem(id int64, product Product, quantity int) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	return &OrderItem{
		ID:        id,
		ProductID: product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  quantity,
	}, nil
}
