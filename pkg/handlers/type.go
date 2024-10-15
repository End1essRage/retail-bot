package handlers

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuHandler interface {
	Menu(u *bot.TgRequest)
	CategorySelect(c *bot.TgRequest)
	ProductSelect(c *bot.TgRequest)
	Back(c *bot.TgRequest)
	Add(c *bot.TgRequest)
}

type CartHandler interface {
	Cart(c *bot.TgRequest)
	Clear(c *bot.TgRequest)
	Increment(c *bot.TgRequest)
	Decrement(c *bot.TgRequest)
}

type BaseHandler struct {
	bot *tgbotapi.BotAPI
	//перевести с апи на сервис
	api      api.Api
	service  *service.Service
	mFactory factories.MenuMurkupFactory
	cFactory factories.CartMurkupFactory
}

func NewBaseHandler(bot *tgbotapi.BotAPI, api api.Api, service *service.Service) *BaseHandler {
	factory := factories.NewUserMurkupFactory()
	return &BaseHandler{bot: bot, api: api, service: service, mFactory: factory, cFactory: factory}
}
