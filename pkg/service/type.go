package service

import c "github.com/end1essrage/retail-bot/pkg"

type Cart struct {
	Positions []c.Position
	UserName  string
}

func NewCart(userName string) *Cart {
	positions := make([]c.Position, 0)
	return &Cart{Positions: positions, UserName: userName}
}
