package handlers

import (
	"strconv"

	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/factories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *BaseHandler) Cart(c *bot.TgRequest) {
	cart := h.service.GetCart(c.Upd.Message.From.UserName)

	msg := tgbotapi.NewMessage(c.Upd.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = h.cFactory.CreateCartMenu(cart.Positions)

	h.bot.Send(msg)
}

func (h *BaseHandler) Clear(c *bot.TgRequest) {
	h.service.ClearCart(c.Upd.CallbackQuery.From.UserName)

	msg := h.formatRootMenu(c.Upd.Message.Chat.ID)

	h.bot.Send(msg)
}

func (h *BaseHandler) Increment(c *bot.TgRequest) {
	msg := h.changeAmount(c, 1)
	h.bot.Send(msg)
}

func (h *BaseHandler) Decrement(c *bot.TgRequest) {
	msg := h.changeAmount(c, -1)
	h.bot.Send(msg)
}

func (h *BaseHandler) changeAmount(c *bot.TgRequest, amount int) tgbotapi.MessageConfig {
	productId, err := strconv.Atoi(c.Data.Data[factories.Product_Id])
	if err != nil {
		logrus.Error("error")
	}

	cart := h.service.ChangeProductAmountInCart(c.Upd.Message.From.UserName, productId, amount)

	msg := tgbotapi.NewMessage(c.Upd.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = h.cFactory.CreateCartMenu(cart.Positions)

	return msg
}
