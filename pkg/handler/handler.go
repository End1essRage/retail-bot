package handler

import (
	"strconv"
	"strings"

	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TgHandler struct {
	bot *tgbotapi.BotAPI
	api api.IApi
}

func NewTgHandler(bot *tgbotapi.BotAPI, api api.IApi) *TgHandler {
	return &TgHandler{bot: bot, api: api}
}

func (h *TgHandler) Handle(u *tgbotapi.Update) {

	if u.Message != nil {
		var reply tgbotapi.MessageConfig

		if u.Message.IsCommand() {
			switch u.Message.Command() {
			case "start":
				reply = h.handleStart(u)
			case "menu":
				reply = h.handleMenu(u)
			default:
				reply = tgbotapi.NewMessage(u.Message.Chat.ID, "Unknown Command")
			}
		}

		h.bot.Send(reply)

	} else if u.CallbackQuery != nil {
		callback := u.CallbackQuery

		data := callback.Data

		shards := strings.Split(data, "_")
		shards = shards[:len(shards)-1]
		command := strings.Join(shards, "_")

		deleteMsg := tgbotapi.NewDeleteMessage(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID)

		h.bot.Send(deleteMsg)

		switch command {
		case "c_id":
			h.bot.Send(h.handleCategoryNavigation(callback))
		case "p_id":
			confirmationMessage := tgbotapi.NewMessage(callback.Message.Chat.ID, "Вы нажали: "+data)

			h.bot.Send(confirmationMessage)
		}

		logrus.Printf("Кнопка нажата: %s", data)
	}
}

func (h *TgHandler) handleStart(u *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.Message.Chat.ID, "hello")
}

func (h *TgHandler) handleMenu(u *tgbotapi.Update) tgbotapi.MessageConfig {
	categories, err := h.api.GetCategories()
	if err != nil {
		return tgbotapi.NewMessage(u.Message.Chat.ID, "error: "+err.Error())
	}

	if len(categories) < 1 {
		return tgbotapi.NewMessage(u.Message.Chat.ID, "error: No categories")
	}

	//Создаем объекты кнопок
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(c.Name, "c_id_"+strconv.Itoa(c.Id)))
	}

	//Создаем двумерный слайс по параметрам

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 3)

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите Категорию:")
	msg.ReplyMarkup = inlineKeyboard

	return msg
}

func (h *TgHandler) handleCategoryNavigation(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	cbData := strings.Split(c.Data, "_")
	categoryId, err := strconv.Atoi(cbData[len(cbData)-1])
	if err != nil {
		logrus.Error(err)
	}

	products, err := h.api.GetProducts(categoryId)
	if err != nil {
		return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
	}
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range products {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(p.Name, "p_id_"+strconv.Itoa(p.Id)))
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 2)

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите товар:")
	msg.ReplyMarkup = inlineKeyboard

	return msg
}

func groupButtons(buttons []tgbotapi.InlineKeyboardButton, inRow int) [][]tgbotapi.InlineKeyboardButton {
	originalSlice := buttons

	var result [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < len(originalSlice); i += inRow {
		end := i + inRow

		if end > len(originalSlice) {
			end = len(originalSlice)
		}
		result = append(result, originalSlice[i:end])
	}

	return result
}
