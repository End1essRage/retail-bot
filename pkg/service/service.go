package service

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/patrickmn/go-cache"
)

// сделать интерфейс
// расщирить кэш для хранения id чата админов и список админов

type Service struct {
	api   api.Api
	cache *cache.Cache
}

func NewServie(api api.Api, cache *cache.Cache) *Service {
	return &Service{api: api, cache: cache}
}
