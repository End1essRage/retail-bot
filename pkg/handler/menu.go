package handler

import (
	"os"
	"time"

	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *TgHandler) handleMenu(u *tgbotapi.Update) tgbotapi.MessageConfig {
	//Запрос категорий с сервера
	categories := h.loadCategories()

	if len(categories) < 1 {
		return tgbotapi.NewMessage(u.Message.Chat.ID, "error: No categories")
	}

	categoriesFiltered := make([]api.Category, 0)
	for _, cat := range categories {
		if cat.Parent == 0 {
			categoriesFiltered = append(categoriesFiltered, cat)
		}
	}
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Выберите Категорию:")
	msg.ReplyMarkup = h.mFactory.CreateRootMenu(categoriesFiltered)

	return msg
}

func (h *TgHandler) handleCategorySelect(c *tgbotapi.CallbackQuery, categoryId int) tgbotapi.MessageConfig {
	currentId := categoryId
	h.cache.Set("currentId", categoryId, 5*time.Minute)

	//подгружаем меню
	menu := h.getMenu()

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
		msg.ReplyMarkup = h.mFactory.CreateProductSelectMenu(categoryId, products)

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
		msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(categories)

		return msg
	}
}

func (h *TgHandler) handleProductSelect(c *tgbotapi.CallbackQuery, productId int) tgbotapi.MessageConfig {
	product, err := h.api.GetProduct(productId)
	if err != nil {
		return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
	}

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, product.Name+"\n"+product.Description)
	msg.ReplyMarkup = h.mFactory.CreateProductMenu(product)

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

func (h *TgHandler) handleBack(c *tgbotapi.CallbackQuery, currentId int, isInProduct bool) tgbotapi.MessageConfig {
	menu := h.getMenu()
	childs := make([]int, 0)
	for _, cat := range menu {
		if cat.Parent == currentId {
			childs = append(childs, cat.Id)
		}
	}

	parentId := h.getParent(currentId)
	//залогировать таргеты и сделать ветвление
	if isInProduct {
		products, err := h.api.GetProducts(currentId)
		if err != nil {
			return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
		}

		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите товар:")
		msg.ReplyMarkup = h.mFactory.CreateProductSelectMenu(currentId, products) // parentId

		return msg
	}
	if parentId == 0 {
		categories := h.getMenu()
		categoriesFiltered := make([]api.Category, 0)
		for _, cat := range categories {
			if cat.Parent == 0 {
				categoriesFiltered = append(categoriesFiltered, cat)
			}
		}
		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите Категорию:")
		msg.ReplyMarkup = h.mFactory.CreateRootMenu(categoriesFiltered)

		return msg
	} else {
		categories := h.getChilds(parentId)

		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите Категорию:")
		msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(categories)

		return msg
	}
}

func (h *TgHandler) getChilds(id int) []api.Category {
	menu := h.getMenu()
	result := make([]api.Category, 0)
	for _, c := range menu {
		if c.Parent == id {
			result = append(result, c)
		}
	}
	return result
}

func (h *TgHandler) loadCategories() []api.Category {
	categories, err := h.api.GetCategories()
	if err != nil {
		logrus.Error(err.Error())
	}
	h.cache.Add("menu", categories, 5*time.Minute)

	return categories
}

func (h *TgHandler) getParent(curId int) int {
	parentId := 0
	menu := h.getMenu()
	for _, item := range menu {
		if item.Id == curId {
			parentId = item.Parent
		}
	}
	return parentId
}

func (h *TgHandler) getMenu() []api.Category {
	var menu []api.Category
	item, ok := h.cache.Get("menu")
	if ok {
		menu = item.([]api.Category)
	} else {
		menu = h.loadCategories()
	}
	return menu
}