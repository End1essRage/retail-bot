package service

import (
	"fmt"
	"strings"

	"github.com/end1essrage/retail-bot/pkg/api"
	f "github.com/end1essrage/retail-bot/pkg/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func (s *Service) CreateOrder(chatId int64, userName string) ([]tgbotapi.MessageConfig, error) {
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
	orderId, err := s.api.CreateOrder(request)
	if err != nil || orderId < 0 {
		return msgs, fmt.Errorf("ошибка создания заказа")
	}

	//сформировать сообщение с составом заказа и кнопками принять и отменить
	sb := strings.Builder{}
	sb.WriteString("Новый заказ от " + userName + "\n")
	for _, p := range cart.Positions {
		sb.WriteString(p.String())
	}

	//admin informing
	msg := tgbotapi.NewMessage(viper.GetInt64("admin_chat_id"), sb.String())

	markup := f.CreateOrderManagerButtonGroup(chatId, orderId)
	msg.ReplyMarkup = markup

	msgs = append(msgs, msg)

	return msgs, nil
}
