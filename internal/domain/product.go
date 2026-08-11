package domain

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
}

func NewProduct(id int64, name string, desc string, price float64) (*Product, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	return &Product{
		ID:          id,
		Name:        name,
		Description: desc,
		Price:       price,
	}, nil
}
