package bot

import (
	"github.com/end1essrage/retail-bot/pkg/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	commandHandlers  map[string]CommandHandler
	callbackHandlers map[string]CallbackHandler
	middlewares      []Middleware
}

func NewBot() *Bot {
	return &Bot{commandHandlers: make(map[string]CommandHandler), callbackHandlers: make(map[string]CallbackHandler)}
}

type TgRequest struct {
	Upd  *tgbotapi.Update
	Data helpers.CallbackData
}

type CommandHandler func(req *TgRequest)

type CallbackHandler func(req *TgRequest)

type Middleware func(req *TgRequest, next Handler) Handler

type Handler func(req *TgRequest)
