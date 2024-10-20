package handlers

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot *tgbotapi.BotAPI
	//перевести с апи на сервис
	api      *api.Api
	service  *service.Service
	mFactory *factories.MurkupFactory
}

func NewHandler(bot *tgbotapi.BotAPI, api *api.Api, service *service.Service) *Handler {
	factory := factories.NewMurkupFactory()
	return &Handler{bot: bot, api: api, service: service, mFactory: factory}
}
