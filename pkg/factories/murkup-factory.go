package factories

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateRootMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		bt := createCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func CreateCategorySelectMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	parentId := 0

	for _, c := range categories {
		bt := createCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
		parentId = c.Parent
	}

	buttons = append(buttons, createBackButton(parentId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func CreateProductSelectMenu(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range Products {
		buttons = append(buttons, createProductSelectButton(p.Name, p.Id))
	}

	buttons = append(buttons, createBackButton(categoryId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func CreateProductMenu(Product api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, createAddButton(Product.Id, Product.Name))
	buttons = append(buttons, createBackButton(Product.CategoryId, true))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func CreateCartMenu(positions []service.Position) tgbotapi.InlineKeyboardMarkup {
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)

	for _, pos := range positions {
		positionButtons := createPositionButtonGroup(pos.Product.Id, pos.Product.Name, pos.Count)
		buttons = append(buttons, positionButtons[0])
		buttons = append(buttons, positionButtons[1])
	}

	navButtons := make([]tgbotapi.InlineKeyboardButton, 0)
	navButtons = append(navButtons, createClearCartButton())
	navButtons = append(navButtons, createOrderCreationButton())

	buttons = append(buttons, navButtons)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = buttons

	return inlineKeyboard
}

func CreateOrdersListMenu(orders []api.OrderShort) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	for _, order := range orders {
		inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, createOrderShortButtonGroup(order))
	}

	return inlineKeyboard
}

func CreateOrderInfo(order api.Order) tgbotapi.InlineKeyboardMarkup {
	buttons := createOrderButtonGroup(order.Id)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, buttons)

	return inlineKeyboard
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
