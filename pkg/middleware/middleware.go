package middleware

import (
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/helpers"
)

type Middleware func(req *bot.TgRequest, next bot.Handler) bot.Handler

func CallbackDataExtruderMiddleware(req *bot.TgRequest, next bot.Handler) bot.Handler {
	return func(update *bot.TgRequest) {
		if req.Upd.CallbackQuery == nil {
			next(update)
			return
		}

		data := helpers.GetCallBackTypeAndData(update.Upd.CallbackQuery)

		req.Data = *data

		next(update)
	}
}
