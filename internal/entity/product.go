package entity

type Product struct {
	id    int
	name  string
	price int
}

func NewProduct(id int, name string, price int) Product {
	return Product{
		id:    id,
		name:  name,
		price: price,
	}
}

func (p *Product) GetId() int {
	return p.id
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() int {
	return p.price
}
