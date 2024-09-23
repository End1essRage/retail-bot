package handler

import (
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

// refactor
func (h *TgHandler) handleAdd(c *tgbotapi.CallbackQuery, productId int, productName string) tgbotapi.MessageConfig {
	h.service.AddProductToCart(c.From.UserName, service.NewProduct(productId, productName))

	msg := h.formatRootMenu(c.Message.Chat.ID)

	return msg
}

func (h *TgHandler) handleClearCart(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	h.service.ClearCart(c.From.UserName)

	msg := h.formatRootMenu(c.Message.Chat.ID)

	return msg
}

func (h *TgHandler) handleCreateOrder(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	if err := h.service.CreateOrder(c.From.UserName); err != nil {
		logrus.Error(err)
	}
	return tgbotapi.NewMessage(c.Message.Chat.ID, "NOT IMPLEMENTED")
}

// dubles
func (h *TgHandler) handleIncrement(c *tgbotapi.CallbackQuery, productId int) tgbotapi.MessageConfig {
	cart := h.service.ChangeProductAmountInCart(c.From.UserName, productId, 1)

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = h.mFactory.CreateCartMenu(cart.Positions)

	return msg
}

func (h *TgHandler) handleDecrement(c *tgbotapi.CallbackQuery, productId int) tgbotapi.MessageConfig {
	cart := h.service.ChangeProductAmountInCart(c.From.UserName, productId, -1)

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, "cart is :")
	msg.ReplyMarkup = h.mFactory.CreateCartMenu(cart.Positions)

	return msg
}
