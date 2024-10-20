package service

import (
	"fmt"

	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func (s *Service) CreateOrder(userName string) ([]tgbotapi.MessageConfig, error) {
	msgs := make([]tgbotapi.MessageConfig, 0)
	cart := s.GetCart(userName)

	if len(cart.Positions) < 1 {
		return msgs, fmt.Errorf("cart is empty")
	}

	request := api.CreateOrderRequest{UserName: userName}
	items := make(map[int]int)
	for _, pos := range cart.Positions {
		items[pos.Product.Id] = pos.Count
	}

	request.Positions = items
	if err := s.api.CreateOrder(request); err != nil {
		return msgs, fmt.Errorf("Ошибка создания заказа")
	}

	//сформировать сообщение с составом заказа и кнопками принять и отменить
	msg := tgbotapi.NewMessage(viper.GetInt64("admin_chat_id"), "Новый заказ")
	msgs = append(msgs, msg)

	return msgs, nil
}

func (s *Service) ChangeOrderStatus() ([]tgbotapi.MessageConfig, error) {
	msgs := make([]tgbotapi.MessageConfig, 0)

	//inform admin chat, user

	return msgs, nil
}
