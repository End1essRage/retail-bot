package handlers

import (
	"strconv"
	"strings"

	cons "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/bot"
	f "github.com/end1essrage/retail-bot/pkg/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Orders(c *bot.TgRequest) {
	username := c.UserName
	orders, err := h.api.GetOrders(username)
	if err != nil {
		logrus.Error(err.Error())
	}

	mu := f.CreateOrdersListMenu(orders)
	msg := tgbotapi.NewMessage(c.ChatId, "your orders: ")
	msg.ReplyMarkup = mu

	h.bot.Send(msg)
}

func (h *Handler) OrderInfo(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	orderId, err := strconv.Atoi(c.Data.Data[f.Order_Id])
	if err != nil {
		logrus.Error(err.Error())
	}

	order, err := h.api.GetOrder(orderId)
	if err != nil {
		logrus.Error(err.Error())
	}
	markup := f.CreateOrderInfo(order)
	//запролнить сообщение с составом заказа
	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, h.formatPositionsString(order.Positions))
	//добавить кнопку отменить и кнопку назад
	msg.ReplyMarkup = markup

	h.bot.Send(msg)
}

func (h *Handler) formatPositionsString(items []api.Position) string {
	sb := strings.Builder{}
	sb.WriteString("Состав заказа : \n")
	for _, item := range items {
		sb.WriteString(item.String() + "\n")
	}

	return sb.String()
}

func (h *Handler) OrderBack(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)
	h.Orders(c)
}

func (h *Handler) CancelOrder(c *bot.TgRequest) {
	//сообщение с  подтверждением отмены

	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	orderId, err := strconv.Atoi(c.Data.Data[f.Order_Id])
	if err != nil {
		logrus.Error(err.Error())
	}

	if err := h.api.ChangeOrderStatus(orderId, int(cons.Cancelled)); err != nil {
		logrus.Error(err.Error())
	}
	//inform user or admin
	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Заказ отменен")

	h.bot.Send(msg)
}

func (h *Handler) AcceptOrder(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	orderId, err := strconv.Atoi(c.Data.Data[f.Order_Id])
	if err != nil {
		logrus.Error(err.Error())
	}

	if err := h.api.ChangeOrderStatus(orderId, int(cons.Accepted)); err != nil {
		logrus.Error(err.Error())
	}
	//inform user
	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Заказ принят")

	h.bot.Send(msg)
}
