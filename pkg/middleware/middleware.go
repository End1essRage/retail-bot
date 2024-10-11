package middleware

import (
	"github.com/end1essrage/retail-bot/pkg/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Middleware func(next bot.Handler) bot.Handler

func adminMiddleware(next bot.Handler) bot.Handler {
	return func(update tgbotapi.Update) {
		if true {
			next(update) // Вызываем следующий хендлер
		} else {
			// Отправляем сообщение неадминистратору

			return // Прекращаем обработку
		}
	}
}
