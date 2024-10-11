package bot

import (
	"github.com/end1essrage/retail-bot/pkg/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *Bot) RegisterCommand(command string, handler CommandHandler) {
	b.commandHandlers[command] = handler
}

func (b *Bot) RegisterCallback(callback string, handler CallbackHandler) {
	b.callbackHandlers[callback] = handler
}

func (b *Bot) Use(mw Middleware) {
	b.middlewares = append(b.middlewares, mw)
}

func (b *Bot) HandleUpdate(update tgbotapi.Update) {
	var handler Handler = func(update tgbotapi.Update) {
		// Обработка команд или колбеков
		if update.Message != nil {
			// Команды
			if handler, exists := b.commandHandlers[update.Message.Command()]; exists {
				handler(update)
			}
		} else if update.CallbackQuery != nil {
			// Обработка колбеков
			data, err := helpers.GetCallBackTypeAndData(update.CallbackQuery)
			if err != nil {
				logrus.Error("error handling")
			}
			if handler, exists := b.callbackHandlers[string(data.Type)]; exists {
				handler(update)
			}
		}
	}
	// Перебираем мидлвары в обратном порядке
	for i := len(b.middlewares) - 1; i >= 0; i-- {
		handler = b.middlewares[i](handler)
	}
	// Запускаем цепочку
	handler(update)
}
