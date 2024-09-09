package handler

import (
	"os"
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TgHandler struct {
	bot      *tgbotapi.BotAPI
	api      api.IApi
	bFactory factories.ButtonsFactory
	mFactory *factories.MurkupFactory
}

func NewTgHandler(bot *tgbotapi.BotAPI, api api.IApi, bfactory factories.ButtonsFactory, mfactory *factories.MurkupFactory) *TgHandler {
	return &TgHandler{bot: bot, api: api, bFactory: bfactory, mFactory: mfactory}
}

func (h *TgHandler) Handle(u *tgbotapi.Update) {

	if u.Message != nil {
		//HAndling commands
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
		//handling buttons
		callback := u.CallbackQuery
		data, err := helpers.GetCallBackTypeAndData(callback)
		if err != nil {
			h.bot.Send(h.SendError(callback, err.Error()))
		}

		//удаление старого сообщения
		deleteMsg := tgbotapi.NewDeleteMessage(u.CallbackQuery.Message.Chat.ID, u.CallbackQuery.Message.MessageID)
		h.bot.Send(deleteMsg)

		switch data.Type {
		case c.CategorySelect:
			categoryId, err := strconv.Atoi(data.Data[factories.CategorySelect_Id])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}

			h.bot.Send(h.handleCategorySelect(callback, categoryId))
		case c.ProductSelect:
			productId, err := strconv.Atoi(data.Data[factories.CategorySelect_Id])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}

			h.bot.Send(h.handleProductSelect(callback, productId))
		case c.Back:
			h.bot.Send(h.handleBack(callback))
		}

		logrus.Printf("Кнопка нажата: %s", data)
	}
}

func (h *TgHandler) handleStart(u *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.Message.Chat.ID, "hello")
}

func (h *TgHandler) handleMenu(u *tgbotapi.Update) tgbotapi.MessageConfig {
	//Запрос категорий с сервера
	categories, err := h.api.GetCategories()
	if err != nil {
		return tgbotapi.NewMessage(u.Message.Chat.ID, "error: "+err.Error())
	}
	if len(categories) < 1 {
		return tgbotapi.NewMessage(u.Message.Chat.ID, "error: No categories")
	}

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите Категорию:")
	msg.ReplyMarkup = h.mFactory.CreateRootMenu(categories)

	return msg
}

func (h *TgHandler) handleCategorySelect(c *tgbotapi.CallbackQuery, categoryId int) tgbotapi.MessageConfig {
	products, err := h.api.GetProducts(categoryId)
	if err != nil {
		return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
	}
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range products {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(p.Name, "p_id_"+strconv.Itoa(p.Id)))
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите товар:")
	msg.ReplyMarkup = inlineKeyboard

	return msg
}

func (h *TgHandler) handleProductSelect(c *tgbotapi.CallbackQuery, productId int) tgbotapi.MessageConfig {
	product, err := h.api.GetProduct(productId)
	if err != nil {
		return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
	}

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Add", "p_add_"+strconv.Itoa(product.Id)))
	buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Back", "p_back_0"))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 2)

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, product.Name+"\n"+product.Description)
	msg.ReplyMarkup = inlineKeyboard

	return msg
}

func (h *TgHandler) handleAdd(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	photoBytes, err := os.ReadFile("/home/end1essrage/Projects/retail-bot/files/memi-klev-club-shai-p-memi-negr-na-krovati-6.jpg")
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}
	imageMessage := tgbotapi.NewPhoto(c.Message.Chat.ID, photoFileBytes)

	h.bot.Send(imageMessage)

	return tgbotapi.NewMessage(c.Message.Chat.ID, ")))))   /menu")
}

func (h *TgHandler) handleBack(c *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {

	//Создаем двумерный слайс по параметрам

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите Категорию:")
	msg.ReplyMarkup = inlineKeyboard

	return msg
}

func (h *TgHandler) SendError(c *tgbotapi.CallbackQuery, err string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(c.Message.Chat.ID, err)
}

// refactor
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
