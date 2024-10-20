package router

import (
	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/handlers"
	"github.com/end1essrage/retail-bot/pkg/middleware"
)

func MapHandlers(handler *handlers.Handler) *bot.Bot {
	bot := bot.NewBot()

	bot.RegisterCommand("menu", handler.Menu)
	bot.RegisterCommand("cart", handler.Cart)
	bot.RegisterCommand("orders", handler.Orders)

	bot.RegisterCallback(string(c.ProductSelect), handler.ProductSelect)
	bot.RegisterCallback(string(c.CategorySelect), handler.CategorySelect)
	bot.RegisterCallback(string(c.Back), handler.Back)

	bot.RegisterCallback(string(c.ProductAdd), handler.Add)
	bot.RegisterCallback(string(c.ClearCart), handler.Clear)
	bot.RegisterCallback(string(c.ProductIncrement), handler.Increment)
	bot.RegisterCallback(string(c.ProductDecrement), handler.Decrement)
	bot.RegisterCallback(string(c.CreateOrder), handler.CreateOrder)

	bot.RegisterCallback(string(c.OrderShortOpen), handler.OrderInfo)
	bot.RegisterCallback(string(c.OrderBackToList), handler.OrderBack)
	bot.RegisterCallback(string(c.OrderChangeStatus), handler.ChangeOrderStatus)

	//вытаскивает данные из колбеков
	bot.Use(middleware.CallbackDataExtruderMiddleware)

	return bot
}
