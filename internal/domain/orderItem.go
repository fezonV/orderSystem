package domain

type OrderItem struct {
	ProductID int64
	Name      string
	Price     float64
	Quantity  int
}

func NewOrderItem(product Product, quantity int) (*OrderItem, error) {

	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	return &OrderItem{
		ProductID: product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  quantity,
	}, nil
}
