package handlers

import (
	"strconv"

	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/bot"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/helpers"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *BaseHandler) Menu(u *bot.TgRequest) {
	//Запрос категорий с сервера
	logrus.Info("ьутг")
	msg := h.formatRootMenu(u.Upd.Message.Chat.ID)
	h.bot.Send(msg)
}

func (h *BaseHandler) Add(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	productId, err := strconv.Atoi(c.Data.Data[factories.Product_Id])
	if err != nil {
		logrus.Error("error")
	}
	productName := c.Data.Data[factories.Product_Name]

	h.service.AddProductToCart(c.Upd.CallbackQuery.From.UserName, service.NewProduct(productId, productName))

	msg := h.formatRootMenu(c.Upd.CallbackQuery.Message.Chat.ID)

	h.bot.Send(msg)
}

func (h *BaseHandler) CategorySelect(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	categoryId, err := strconv.Atoi(c.Data.Data[factories.Category_Id])
	if err != nil {
		logrus.Error("Ошибка считывания параметра")
	}

	currentId := categoryId

	//подгружаем меню
	menu := h.service.GetMenu()

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
		logrus.Info("last")
		products, err := h.api.GetProducts(categoryId)
		if err != nil {
			h.bot.Send(tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "error: "+err.Error()))
			return
		}

		msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Выберите товар:")
		msg.ReplyMarkup = h.mFactory.CreateProductSelectMenu(categoryId, products)

		h.bot.Send(msg)
		return
	} else {
		logrus.Info("not last")
		categories := make([]api.Category, 0)
		for _, i := range childs {
			for _, c := range menu {
				if c.Id == i {
					categories = append(categories, c)
				}
			}
		}

		msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Выберите Категорию:")
		msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(categories)

		h.bot.Send(msg)
		return
	}
}

func (h *BaseHandler) ProductSelect(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	productId, err := strconv.Atoi(c.Data.Data[factories.Product_Id])
	if err != nil {
		logrus.Info("error reading data")
	}
	product, err := h.api.GetProductData(productId)
	if err != nil {
		logrus.Info("error reading data")
	}

	logrus.Info(productId)

	msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, product.Name+"\n"+product.Description)
	msg.ReplyMarkup = h.mFactory.CreateProductMenu(product)

	h.bot.Send(msg)
}

func (h *BaseHandler) Back(c *bot.TgRequest) {
	h.deleteMessage(c.Upd.CallbackQuery.Message.Chat.ID, c.Upd.CallbackQuery.Message.MessageID)

	currentId, err := strconv.Atoi(c.Data.Data[factories.Back_CurrentId])
	if err != nil {
		logrus.Info("error reading data")
	}

	isInProduct, err := strconv.ParseBool(c.Data.Data[factories.Back_IsProduct])
	if err != nil {
		logrus.Info("error reading data")
	}

	menu := h.service.GetMenu()
	childs := make([]int, 0)
	for _, cat := range menu {
		if cat.Parent == currentId {
			childs = append(childs, cat.Id)
		}
	}

	parentId := h.service.GetParent(currentId)
	//залогировать таргеты и сделать ветвление
	if isInProduct {
		products, err := h.api.GetProducts(currentId)
		if err != nil {
			logrus.Info("error getting products")
		}

		msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Выберите товар:")
		msg.ReplyMarkup = h.mFactory.CreateProductSelectMenu(currentId, products) // parentId

		h.bot.Send(msg)
	} else {

		if parentId == 0 {

			categories := h.service.GetMenu()
			categoriesFiltered := make([]api.Category, 0)
			for _, cat := range categories {
				if cat.Parent == 0 {
					categoriesFiltered = append(categoriesFiltered, cat)
				}
			}
			msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Выберите Категорию:")
			msg.ReplyMarkup = h.mFactory.CreateRootMenu(categoriesFiltered)

			h.bot.Send(msg)
		} else {
			categories := h.service.GetChilds(parentId)

			msg := tgbotapi.NewMessage(c.Upd.CallbackQuery.Message.Chat.ID, "Выберите Категорию:")
			msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(categories)

			h.bot.Send(msg)
		}
	}
}

func (h *BaseHandler) formatRootMenu(chatId int64) tgbotapi.MessageConfig {
	categories := h.service.GetMenu()
	if len(categories) < 1 {
		return tgbotapi.NewMessage(chatId, "error: No categories")
	}

	categoriesFiltered := helpers.FilterRootCategories(categories)

	msg := tgbotapi.NewMessage(chatId, "Выберите Категорию:")
	msg.ReplyMarkup = h.mFactory.CreateRootMenu(categoriesFiltered)

	return msg
}

func (h *BaseHandler) deleteMessage(chatId int64, messageId int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatId, messageId)
	h.bot.Send(deleteMsg)
}
