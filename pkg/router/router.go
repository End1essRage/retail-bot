package router

import (
	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/handlers"
	"github.com/end1essrage/retail-bot/pkg/middleware"
)

func MapHandlers(mHandler handlers.NavigationHandler, cHandler handlers.CartHandler, oHandler handlers.OrderHandler) *bot.Bot {
	bot := bot.NewBot()

	bot.RegisterCommand("menu", mHandler.Menu)
	bot.RegisterCommand("cart", cHandler.Cart)
	bot.RegisterCommand("orders", oHandler.Orders)

	bot.RegisterCallback(string(c.ProductSelect), mHandler.ProductSelect)
	bot.RegisterCallback(string(c.CategorySelect), mHandler.CategorySelect)
	bot.RegisterCallback(string(c.Back), mHandler.Back)

	bot.RegisterCallback(string(c.ProductAdd), cHandler.Add)
	bot.RegisterCallback(string(c.ClearCart), cHandler.Clear)
	bot.RegisterCallback(string(c.ProductIncrement), cHandler.Increment)
	bot.RegisterCallback(string(c.ProductDecrement), cHandler.Decrement)
	bot.RegisterCallback(string(c.CreateOrder), cHandler.CreateOrder)

	//вытаскивает данные из колбеков
	bot.Use(middleware.CallbackDataExtruderMiddleware)

	return bot
}
