package handler

import (
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *TgHandler) handleAdd(c *tgbotapi.CallbackQuery, productId int, productName string) tgbotapi.MessageConfig {
	h.service.AddProductToCart(c.From.UserName, service.NewProduct(productId, productName))

	return tgbotapi.NewMessage(c.Message.Chat.ID, "Успешно добавлено в корзину /cart")
}
