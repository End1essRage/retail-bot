package constants

import "strconv"

type Product struct {
	Id   int
	Name string
}

func NewProduct(id int, name string) Product {
	return Product{Id: id, Name: name}
}

type Position struct {
	Product Product
	Count   int
}

func (p Position) String() string {
	return p.Product.Name + " X " + strconv.Itoa(p.Count)
}
