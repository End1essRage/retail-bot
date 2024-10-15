package handlers

import (
	"strconv"

	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *BaseHandler) CreateOrder(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	if err := h.service.CreateOrder(c.Upd.CallbackQuery.From.UserName); err != nil {
		logrus.Error(err.Error())
	}

	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Ваш заказ успешно принят")
	//inform managers
	h.bot.Send(msg)
}

func (h *BaseHandler) Add(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	productId, err := strconv.Atoi(c.Data.Data[factories.Product_Id])
	if err != nil {
		logrus.Error("error")
	}
	productName := c.Data.Data[factories.Product_Name]

	h.service.AddProductToCart(c.Upd.CallbackQuery.From.UserName, service.NewProduct(productId, productName))

	msg := h.formatRootMenu(c.Upd.CallbackQuery.Message.Chat.ID)

	h.bot.Send(msg)
}

func (h *BaseHandler) Cart(c *bot.TgRequest) {
	cart := h.service.GetCart(c.Upd.Message.From.UserName)

	msg := tgbotapi.NewMessage(c.Upd.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = h.cFactory.CreateCartMenu(cart.Positions)

	h.bot.Send(msg)
}

func (h *BaseHandler) Clear(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	h.service.ClearCart(c.Upd.CallbackQuery.From.UserName)

	msg := h.formatRootMenu(c.Upd.Message.Chat.ID)

	h.bot.Send(msg)
}

func (h *BaseHandler) Increment(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	logrus.Warning(c.Upd.CallbackQuery.From.UserName)
	msg := h.changeAmount(c, 1)
	h.bot.Send(msg)
}

func (h *BaseHandler) Decrement(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	msg := h.changeAmount(c, -1)
	h.bot.Send(msg)
}

func (h *BaseHandler) changeAmount(c *bot.TgRequest, amount int) tgbotapi.MessageConfig {
	productId, err := strconv.Atoi(c.Data.Data[factories.Product_Id])
	if err != nil {
		logrus.Error("error")
	}
	logrus.Warning(c.Upd.CallbackQuery.From.UserName)
	cart := h.service.ChangeProductAmountInCart(c.Upd.CallbackQuery.From.UserName, productId, amount)

	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = h.cFactory.CreateCartMenu(cart.Positions)

	return msg
}
