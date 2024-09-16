package service

type Cart struct {
	Positions []Position
	UserName  string
}

func NewCart(userName string) *Cart {
	positions := make([]Position, 0)
	return &Cart{Positions: positions, UserName: userName}
}

type Position struct {
	Product Product
	Count   int
}

type Product struct {
	Id   int
	Name string
}

func NewProduct(id int, name string) Product {
	return Product{Id: id, Name: name}
}
