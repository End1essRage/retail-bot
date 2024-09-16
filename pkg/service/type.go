package service

type Cart struct {
	positions []Position
	userName  string
}

func NewCart(userName string) *Cart {
	positions := make([]Position, 0)
	return &Cart{positions: positions, userName: userName}
}

type Position struct {
	product Product
	count   int
}

type Product struct {
	id   int
	name string
}

func NewProduct(id int, name string) Product {
	return Product{id: id, name: name}
}
