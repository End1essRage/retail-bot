package middleware

import (
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type Middleware func(req *bot.TgRequest, next bot.Handler) bot.Handler

func CallbackDataExtruderMiddleware(req *bot.TgRequest, next bot.Handler) bot.Handler {
	logrus.Info("middle")
	return func(update *bot.TgRequest) {
		if req.Upd.CallbackQuery == nil {
			logrus.Info("no query")
			next(update)
			return
		}

		data, err := helpers.GetCallBackTypeAndData(update.Upd.CallbackQuery)
		if err != nil {
			logrus.Error("extruding data")
		}

		req.Data = *data

		next(update)
	}
}
