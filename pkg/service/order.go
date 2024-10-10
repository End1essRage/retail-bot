package service

import (
	"fmt"

	"github.com/end1essrage/retail-bot/pkg/api"
)

func (s *Service) CreateOrder(userName string) error {
	cart := s.GetCart(userName)

	if len(cart.Positions) < 1 {
		return fmt.Errorf("cart is empty")
	}

	request := api.CreateOrderRequest{UserName: userName}
	items := make(map[int]int)
	for _, pos := range cart.Positions {
		items[pos.Product.Id] = pos.Count
	}

	request.Positions = items
	if err := s.api.CreateOrder(request); err != nil {
		return fmt.Errorf("Ошибка создания заказа")
	}

	//send message to admin

	return nil
}
