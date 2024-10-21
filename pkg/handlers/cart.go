package handlers

import (
	"strconv"

	t "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/bot"
	f "github.com/end1essrage/retail-bot/pkg/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Add(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	productId, err := strconv.Atoi(c.Data.Data[f.Product_Id])
	if err != nil {
		logrus.Error("error")
	}
	productName := h.service.GetProductName(productId)

	h.service.AddProductToCart(c.Upd.CallbackQuery.From.UserName, t.NewProduct(productId, productName))

	msg := h.formatRootMenu(c.Upd.CallbackQuery.Message.Chat.ID)

	h.bot.Send(msg)
}

func (h *Handler) Cart(c *bot.TgRequest) {
	cart := h.service.GetCart(c.Upd.Message.From.UserName)

	msg := tgbotapi.NewMessage(c.Upd.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = f.CreateCartMenu(cart.Positions)

	h.bot.Send(msg)
}

func (h *Handler) Clear(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	h.service.ClearCart(c.Upd.CallbackQuery.From.UserName)

	msg := h.formatRootMenu(c.Upd.Message.Chat.ID)

	h.bot.Send(msg)
}

func (h *Handler) Increment(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	logrus.Warning(c.Upd.CallbackQuery.From.UserName)
	msg := h.changeAmount(c, 1)
	h.bot.Send(msg)
}

func (h *Handler) Decrement(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	msg := h.changeAmount(c, -1)
	h.bot.Send(msg)
}

func (h *Handler) changeAmount(c *bot.TgRequest, amount int) tgbotapi.MessageConfig {
	productId, err := strconv.Atoi(c.Data.Data[f.Product_Id])
	if err != nil {
		logrus.Error("error")
	}
	logrus.Warning(c.Upd.CallbackQuery.From.UserName)
	cart := h.service.ChangeProductAmountInCart(c.Upd.CallbackQuery.From.UserName, productId, amount)

	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = f.CreateCartMenu(cart.Positions)

	return msg
}
