package handler

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patrickmn/go-cache"
)

type TgHandler struct {
	bot      *tgbotapi.BotAPI
	api      api.IApi
	cache    *cache.Cache
	bFactory factories.ButtonsFactory
	mFactory *factories.MurkupFactory
}

func NewTgHandler(bot *tgbotapi.BotAPI, api api.IApi, cache *cache.Cache, bfactory factories.ButtonsFactory, mfactory *factories.MurkupFactory) *TgHandler {
	return &TgHandler{bot: bot, api: api, cache: cache, bFactory: bfactory, mFactory: mfactory}
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
			case "admin":
				reply = h.handleAdmin(u)
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
			categoryId, err := strconv.Atoi(data.Data[factories.Category_Id])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}
			h.bot.Send(h.handleCategorySelect(callback, categoryId))

		case c.ProductSelect:
			productId, err := strconv.Atoi(data.Data[factories.Product_Id])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}

			h.bot.Send(h.handleProductSelect(callback, productId))
		case c.Back:
			currentId, err := strconv.Atoi(data.Data[factories.Back_CurrentId])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}
			isInProduct, err := strconv.ParseBool(data.Data[factories.Back_IsProduct])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}
			h.bot.Send(h.handleBack(callback, currentId, isInProduct))

		case c.ProductAdd:
			product, err := strconv.Atoi(data.Data[factories.Product_Id])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}
			h.bot.Send(h.handleAdd(callback, product))
		}
	}
}

func (h *TgHandler) handleStart(u *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.Message.Chat.ID, "hello")
}

func (h *TgHandler) handleAdmin(u *tgbotapi.Update) tgbotapi.MessageConfig {
	//Check is admin
	//create some kind of session
	//send admin menu
	return tgbotapi.NewMessage(u.Message.Chat.ID, "hello")
}

func (h *TgHandler) SendError(c *tgbotapi.CallbackQuery, err string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(c.Message.Chat.ID, err)
}
