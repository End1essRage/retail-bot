package router

import (
	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/handlers"
	"github.com/end1essrage/retail-bot/pkg/middleware"
)

func MapHandlers(mHandler handlers.MenuHandler) *bot.Bot {
	bot := bot.NewBot()

	bot.RegisterCommand("menu", mHandler.Menu)

	bot.RegisterCallback(string(c.ProductSelect), mHandler.ProductSelect)
	bot.RegisterCallback(string(c.Back), mHandler.Back)
	bot.RegisterCallback(string(c.CategorySelect), mHandler.CategorySelect)
	bot.RegisterCallback(string(c.ProductAdd), mHandler.Add)

	bot.Use(middleware.CallbackDataExtruderMiddleware)

	return bot
}
