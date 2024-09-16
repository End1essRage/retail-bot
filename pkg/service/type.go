package service

import "github.com/end1essrage/retail-bot/pkg/api"

type Cart struct {
	positions []Position
	userName  string
}

func NewCart(userName string) *Cart {
	positions := make([]Position, 0)
	return &Cart{positions: positions, userName: userName}
}

type Position struct {
	product api.Product
	count   int
}
