package service

import (
	"time"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type Service struct {
	api   api.IApi
	cache *cache.Cache
}

func NewServie(api api.IApi, cache *cache.Cache) *Service {
	return &Service{api: api, cache: cache}
}

func (s *Service) AddProductToCart(userName string, product api.Product) {
	logrus.Info("product added to cart")

	cart := s.GetCart(userName)
	cart.positions = append(cart.positions, Position{product: product, count: 1})

	s.updateCart(cart)
}

func (s *Service) updateCart(cart *Cart) {
	s.cache.Set(c.CacheCartUserPrefix+"_"+cart.userName, cart, 5*time.Minute)
}

func (s *Service) GetCart(userName string) *Cart {
	data, ok := s.cache.Get(c.CacheCartUserPrefix + "_" + userName)
	var cart Cart
	if ok {
		cart = data.(Cart)
	} else {
		cart = *NewCart(userName)
	}

	return &cart
}
