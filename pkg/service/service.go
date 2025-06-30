package service

import (
	"time"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

// сделать интерфейс
// расщирить кэш для хранения id чата админов и список админов

type Service struct {
	api   *api.Api
	cache *cache.Cache
}

func NewServie(api *api.Api, cache *cache.Cache) *Service {
	return &Service{api: api, cache: cache}
}

func (s *Service) GetUserRole(userName string) c.UserRole {
	roleId := int(c.Client)
	//add cache

	roleCS, ok := s.cache.Get(string(c.CacheUserRolePrefix) + c.CacheSeparator + userName)
	if ok {
		roleCI, ok := roleCS.(int)
		if ok {
			roleId = roleCI
			return c.UserRole(roleId)
		}
	}

	roleId, err := s.api.GetUserRole(userName)
	if err != nil {
		logrus.Error("Ошибка при получении роли" + err.Error())
		s.cache.Set(string(c.CacheUserRolePrefix)+c.CacheSeparator+userName, int(c.Client), 5*time.Minute)
		return c.Client
	}

	s.cache.Set(string(c.CacheUserRolePrefix)+c.CacheSeparator+userName, roleId, 5*time.Minute)

	return c.UserRole(roleId)
}
