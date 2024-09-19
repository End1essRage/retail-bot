package handler

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// refactor
func (h *TgHandler) handleMenu(u *tgbotapi.Update) tgbotapi.MessageConfig {
	//Запрос категорий с сервера
	msg := h.formatRootMenu(u.Message.Chat.ID)
	return msg
}

func (h *TgHandler) handleCategorySelect(c *tgbotapi.CallbackQuery, categoryId int) tgbotapi.MessageConfig {
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
	product, err := h.api.GetProductData(productId)
	if err != nil {
		return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
	}

	msg := tgbotapi.NewMessage(c.Message.Chat.ID, product.Name+"\n"+product.Description)
	msg.ReplyMarkup = h.mFactory.CreateProductMenu(product)

	return msg
}

func (h *TgHandler) handleBack(c *tgbotapi.CallbackQuery, currentId int, isInProduct bool) tgbotapi.MessageConfig {
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
			return tgbotapi.NewMessage(c.Message.Chat.ID, "error: "+err.Error())
		}

		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите товар:")
		msg.ReplyMarkup = h.mFactory.CreateProductSelectMenu(currentId, products) // parentId

		return msg
	}
	if parentId == 0 {
		categories := h.service.GetMenu()
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
		categories := h.service.GetChilds(parentId)

		msg := tgbotapi.NewMessage(c.Message.Chat.ID, "Выберите Категорию:")
		msg.ReplyMarkup = h.mFactory.CreateCategorySelectMenu(categories)

		return msg
	}
}
