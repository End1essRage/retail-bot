package service

import (
	"time"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

// интерфейс

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

func (s *Service) loadCategories() []api.Category {
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
		menu = s.loadCategories()
	}
	return menu
}

func (s *Service) updateCart(cart Cart) {
	s.cache.Set(c.CacheCartUserPrefix+"_"+cart.UserName, cart, 5*time.Minute)
}

func (s *Service) AddProductToCart(userName string, product Product) {
	logrus.Info("product added to cart")

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
		cart.Positions = append(cart.Positions, Position{Product: product, Count: 1})
	} else {
		cart.Positions[dubleId] = Position{Product: product, Count: count}
	}

	s.updateCart(*cart)
}

func (s *Service) ChangeProductAmountInCart(userName string, productId int, lambda int) Cart {
	cart := s.GetCart(userName)

	var position Position

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
	data, ok := s.cache.Get(c.CacheCartUserPrefix + "_" + userName)
	var cart Cart
	if ok {
		cart = data.(Cart)
	} else {
		cart = *NewCart(userName)
	}

	return &cart
}
