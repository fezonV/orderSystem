package domain

type Product struct {
	id          int64
	name        string
	description string
	price       float64
}

func (p Product) ID() int64 {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Price() float64 {
	return p.price
}

func NewProduct(id int64, name string, desc string, price float64) (*Product, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	return &Product{
		id:          id,
		name:        name,
		description: desc,
		price:       price,
	}, nil
}
