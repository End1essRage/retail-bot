package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	commandHandlers  map[string]CommandHandler
	callbackHandlers map[string]CallbackHandler
	middlewares      []Middleware
}

type CommandHandler func(update tgbotapi.Update)

type CallbackHandler func(update tgbotapi.Update)

type Middleware func(next Handler) Handler

type Handler func(update tgbotapi.Update)
