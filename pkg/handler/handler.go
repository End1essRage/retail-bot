package handler

import (
	"os"
	"strconv"
	"time"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
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
			back, err := strconv.Atoi(data.Data[factories.Back_CurrentId])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}
			h.bot.Send(h.handleBack(callback, back))

		case c.ProductAdd:
			product, err := strconv.Atoi(data.Data[factories.Product_Id])
			if err != nil {
				h.bot.Send(h.SendError(callback, err.Error()))
			}
			h.bot.Send(h.handleAdd(callback, product))
		}

		logrus.Printf("Кнопка нажата: %s", data)
	}
}

func (h *TgHandler) handleStart(u *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.Message.Chat.ID, "hello")
}

func (h *TgHandler) handleMenu(u *tgbotapi.Update) tgbotapi.MessageConfig {
	//Запрос категорий с сервера
	categories := h.loadCategories()

	if len(categories) < 1 {
		return tgbotapi.NewMessage(u.Message.Chat.ID, "error: No categories")
	}

	categoriesFiltered := make([]api.Category, 0)
	for _, cat := range categories {
		logrus.Warn(cat)
		logrus.Warn(cat.Parent)
		if cat.Parent == 0 {
			categoriesFiltered = append(categoriesFiltered, cat)
		}
	}
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите Категорию:")
	msg.ReplyMarkup = h.mFactory.CreateRootMenu(categoriesFiltered)

	return msg
}

func (h *TgHandler) handleCategorySelect(c *tgbotapi.CallbackQuery, categoryId int) tgbotapi.MessageConfig {
	parentIdstr, ok := h.cache.Get("currentId")
	if !ok {
		parentIdstr = 0
	}

	logrus.Info("parentid is ", parentIdstr.(int))
	parentId := parentIdstr.(int)

	currentId := categoryId
	h.cache.Set("currentId", categoryId, 5*time.Minute)

	//подгружаем меню
	var menu []api.Category
	item, ok := h.cache.Get("menu")
	if ok {
		menu = item.([]api.Category)
	} else {
		menu = h.loadCategories()
	}

	//проверяем листок ли текущая категория
	isLast := true
	childs := make([]int, 0)
	for _, cat := range menu {
		if cat.Parent == currentId {
			isLast = false
			childs = append(childs, cat.Id)
		}
	}

	if isLast {
		products, err := h.api.GetProducts(categoryId)
		if err != nil {
			return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
		}

		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите товар:")
		msg.ReplyMarkup = h.mFactory.CreateProductSelectMenu(parentId, products)

		return msg
	} else {
		categories := make([]api.Category, 0)
		for _, i := range childs {
			for _, c := range menu {
				if c.Id == i {
					categories = append(categories, c)
				}
			}
		}

		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите Категорию:")
		msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(parentId, categories)

		return msg
	}

}

func (h *TgHandler) handleProductSelect(c *tgbotapi.CallbackQuery, productId int) tgbotapi.MessageConfig {
	product, err := h.api.GetProduct(productId)
	if err != nil {
		return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
	}

	parentIdstr, ok := h.cache.Get("currentId")
	if !ok {
		parentIdstr = 0
	}

	logrus.Info("parentid is ", parentIdstr.(int))
	parentId := parentIdstr.(int)

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, product.Name+"\n"+product.Description)
	msg.ReplyMarkup = h.mFactory.CreateProductMenu(parentId, product)

	return msg
}

func (h *TgHandler) handleAdd(c *tgbotapi.CallbackQuery, productId int) tgbotapi.MessageConfig {
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

func (h *TgHandler) handleBack(c *tgbotapi.CallbackQuery, targetId int) tgbotapi.MessageConfig {

	//залогировать таргеты и сделать ветвление
	var menu []api.Category
	item, ok := h.cache.Get("menu")
	if ok {
		menu = item.([]api.Category)
	} else {
		menu = h.loadCategories()
	}

	logrus.Warn("targetid is ", targetId)
	//var category api.Category
	categories := make([]api.Category, 0)
	upper := 0
	for _, c := range menu {
		if c.Parent == targetId {
			categories = append(categories, c)
		}
		if c.Id == targetId {
			upper = c.Parent
		}
	}

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите Категорию:")
	msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(upper, categories)

	return msg
}

func (h *TgHandler) SendError(c *tgbotapi.CallbackQuery, err string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(c.Message.Chat.ID, err)
}

func (h *TgHandler) loadCategories() []api.Category {
	logrus.Info("categories loaded")

	categories, err := h.api.GetCategories()
	if err != nil {
		logrus.Error(err.Error())
	}
	h.cache.Add("menu", categories, 5*time.Minute)

	return categories
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
