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

func (s *Service) GetChilds(id int) []api.Category {
	menu := s.GetMenu()
	result := make([]api.Category, 0)
	for _, c := range menu {
		if c.Parent == id {
			result = append(result, c)
		}
	}
	return result
}

func (s *Service) LoadCategories() []api.Category {
	categories, err := s.api.GetCategories()
	if err != nil {
		logrus.Error(err.Error())
	}
	s.cache.Add("menu", categories, 5*time.Minute)

	return categories
}

func (s *Service) GetParent(curId int) int {
	parentId := 0
	menu := s.GetMenu()
	for _, item := range menu {
		if item.Id == curId {
			parentId = item.Parent
		}
	}
	return parentId
}

func (s *Service) GetMenu() []api.Category {
	var menu []api.Category
	item, ok := s.cache.Get("menu")
	if ok {
		menu = item.([]api.Category)
	} else {
		menu = s.LoadCategories()
	}
	return menu
}
func (s *Service) updateCart(cart *Cart) {
	s.cache.Set(c.CacheCartUserPrefix+"_"+cart.userName, cart, 5*time.Minute)
}
func (s *Service) AddProductToCart(userName string, product Product) {
	logrus.Info("product added to cart")

	cart := s.GetCart(userName)
	f := false
	for _, pos := range cart.positions {
		if pos.product.id == product.id {
			pos.count++
			f = true
			break
		}
	}

	//если такой позиции в корзине еще нет
	if !f {
		cart.positions = append(cart.positions, Position{product: product, count: 1})
	}

	s.updateCart(cart)
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
