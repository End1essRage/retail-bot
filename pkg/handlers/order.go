package handlers

import (
	"github.com/end1essrage/retail-bot/pkg/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *BaseHandler) Orders(c *bot.TgRequest) {
	//api call
	username := c.Upd.Message.From.UserName
	orders, err := h.api.GetOrders(username)
	if err != nil {
		logrus.Error(err.Error())
	}

	mu := h.oFactory.CreateOrdersListMenu(orders)
	msg := tgbotapi.NewMessage(c.Upd.Message.Chat.ID, "your orders: ")
	msg.ReplyMarkup = mu

	h.bot.Send(msg)
}

func (h *BaseHandler) OrderInfo(c *bot.TgRequest) {

}

func (h *BaseHandler) CancelOrder(c *bot.TgRequest) {

}

func (h *BaseHandler) CloseOrder(c *bot.TgRequest) {

}
