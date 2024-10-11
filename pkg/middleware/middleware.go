package middleware

import (
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type Middleware func(req *bot.TgRequest, next bot.Handler) bot.Handler

func CallbackDataExtruderMiddleware(req *bot.TgRequest, next bot.Handler) bot.Handler {
	return func(update *bot.TgRequest) {
		if req.Upd.CallbackQuery == nil {
			next(req)
		}
		data, err := helpers.GetCallBackTypeAndData(req.Upd.CallbackQuery)
		if err != nil {
			logrus.Error("error extruding callback data")
			return
		}
		req.Data = *data
	}
}
