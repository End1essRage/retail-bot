package service

import (
	"time"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/sirupsen/logrus"
)

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
	s.cache.Add(c.MenuKey, categories, 5*time.Minute)

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
	item, ok := s.cache.Get(c.MenuKey)
	if ok {
		if len(item.([]api.Category)) > 0 {
			menu = item.([]api.Category)
		} else {
			menu = s.loadCategories()
		}
	} else {
		menu = s.loadCategories()
	}
	return menu
}
