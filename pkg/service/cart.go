package service

import (
	"time"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/sirupsen/logrus"
)

func (s *Service) updateCart(cart Cart) {
	logrus.Info("updating cart")
	s.cache.Set(string(c.CacheCartUserPrefix)+c.CacheSeparator+cart.UserName, cart, 5*time.Minute)
}

func (s *Service) AddProductToCart(userName string, product c.Product) {
	cart := s.GetCart(userName)
	f := false
	dubleId := 0
	count := 0
	for i, pos := range cart.Positions {
		if pos.Product.Id == product.Id {
			dubleId = i
			f = true
			count = pos.Count + 1
			break
		}
	}

	//если такой позиции в корзине еще нет
	if !f {
		cart.Positions = append(cart.Positions, c.Position{Product: product, Count: 1})
	} else {
		cart.Positions[dubleId] = c.Position{Product: product, Count: count}
	}

	s.updateCart(*cart)
}

func (s *Service) ChangeProductAmountInCart(userName string, productId int, lambda int) Cart {
	cart := s.GetCart(userName)

	var position c.Position

	dubleId := 0
	count := 0

	for i, pos := range cart.Positions {
		if pos.Product.Id == productId {
			position = pos
			dubleId = i
			count = pos.Count + lambda
			break
		}
	}

	if count < 1 {
		//перезаписать слайс
		result := append(cart.Positions[:dubleId], cart.Positions[dubleId+1:]...)
		cart.Positions = result
	} else {
		position.Count = count
		cart.Positions[dubleId] = position
	}

	s.updateCart(*cart)
	return *cart
}

func (s *Service) GetCart(userName string) *Cart {
	data, ok := s.cache.Get(string(c.CacheCartUserPrefix) + c.CacheSeparator + userName)
	var cart Cart
	if ok {
		cart = data.(Cart)
	} else {
		cart = *NewCart(userName)
	}
	logrus.Info(cart)
	return &cart
}

func (s *Service) ClearCart(userName string) {
	cart := *NewCart(userName)
	s.updateCart(cart)
}
