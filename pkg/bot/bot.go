package bot

import (
	"github.com/end1essrage/retail-bot/pkg/helpers"
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

func (b *Bot) HandleUpdate(req *TgRequest) {
	var handler Handler = func(req *TgRequest) {
		// Обработка команд или колбеков
		if req.Upd.Message != nil {
			if req.Upd.Message.Command() != "" {
				// Команды
				if handler, exists := b.commandHandlers[req.Upd.Message.Command()]; exists {
					handler(req)
				}
			}
		}
		if req.Upd.CallbackQuery != nil {
			// Обработка колбеков
			data := helpers.GetCallBackTypeAndData(req.Upd.CallbackQuery)

			if handler, exists := b.callbackHandlers[string(data.Type)]; exists {
				handler(req)
			}
		}
	}
	// Перебираем мидлвары в обратном порядке
	for i := len(b.middlewares) - 1; i >= 0; i-- {
		handler = b.middlewares[i](req, handler)
	}
	// Запускаем цепочку
	handler(req)
}
