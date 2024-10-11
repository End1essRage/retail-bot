package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (h *TgHandler) handleCreateOrder(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	if err := h.service.CreateOrder(c.From.UserName); err != nil {
		logrus.Error(err)
	}
	h.informAdmins()
	return tgbotapi.NewMessage(c.Message.Chat.ID, "Ваш заказ принят")
}

func (h *TgHandler) handleGetOrders(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	//isAdmin
	//получить список заказов, админы получают все открытые и новые заказы, юзеры все активные свои заказы
	// затем сделаю более удобное управление историей заказов
	// юзеры имеют кнопку с отменой заказа
	// у админов есть кнопка подтверждения(зависит от статуса) и кнопка отмены
	return tgbotapi.NewMessage(c.Message.Chat.ID, "Ваш заказ принят")
}

func (h *TgHandler) handleGetOrder(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	//isAdmin
	//юзеры как и админы могут из краткого списка перейти в карточку заказа
	// тут полный состав заказа и дял админов доп инфа

	return tgbotapi.NewMessage(c.Message.Chat.ID, "Ваш заказ принят")
}

// принять заказ в обработку
func (h *TgHandler) handleAcceptOrder(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	//isAdmin
	// смена статуса возможно заменится на один универсальный метод с цифровым обозначением таргет статуса
	return tgbotapi.NewMessage(c.Message.Chat.ID, "Ваш заказ принят")
}

// отменить заказ
func (h *TgHandler) handleCancelOrder(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	//isAdmin
	//isOwner
	// смена статуса возможно заменится на один универсальный метод с цифровым обозначением таргет статуса
	return tgbotapi.NewMessage(c.Message.Chat.ID, "Ваш заказ принят")
}

// отметка выполнено заказфотменита
func (h *TgHandler) handleCloseOrder(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	//isAdmin
	//isOwner
	// смена статуса возможно заменится на один универсальный метод с цифровым обозначением таргет статуса
	return tgbotapi.NewMessage(c.Message.Chat.ID, "Ваш заказ принят")
}

// Отправка
func (h *TgHandler) informAdmins() {
	// перенести из конфига в кэш в будущем
	msg := tgbotapi.NewMessage(viper.GetInt64("admin_chat_id"), "New Order")
	h.bot.Send(msg)
}

// проинформировать как админов так и владельца заказа
func (h *TgHandler) informStatusChanged() {
	msg := tgbotapi.NewMessage(viper.GetInt64("admin_chat_id"), "New Order")
	h.bot.Send(msg)
}
