package factories

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MurkupFactory struct {
	bFactory *ButtonsFactory
}

func NewMurkupFactory() *MurkupFactory {
	bfactory := NewButtonsFactory()
	return &MurkupFactory{bFactory: bfactory}
}

func (f *MurkupFactory) CreateRootMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		bt := f.bFactory.CreateCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *MurkupFactory) CreateCategorySelectMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	parentId := 0

	for _, c := range categories {
		bt := f.bFactory.CreateCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
		parentId = c.Parent
	}

	buttons = append(buttons, f.bFactory.CreateBackButton(parentId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *MurkupFactory) CreateProductSelectMenu(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range Products {
		buttons = append(buttons, f.bFactory.CreateProductSelectButton(p.Name, p.Id))
	}

	buttons = append(buttons, f.bFactory.CreateBackButton(categoryId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *MurkupFactory) CreateProductMenu(Product api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, f.bFactory.CreateAddButton(Product.Id, Product.Name))
	buttons = append(buttons, f.bFactory.CreateBackButton(Product.CategoryId, true))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *MurkupFactory) CreateCartMenu(positions []service.Position) tgbotapi.InlineKeyboardMarkup {
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)

	for _, pos := range positions {
		positionButtons := f.bFactory.CreatePositionButtonGroup(pos.Product.Id, pos.Product.Name, pos.Count)
		buttons = append(buttons, positionButtons[0])
		buttons = append(buttons, positionButtons[1])
	}

	navButtons := make([]tgbotapi.InlineKeyboardButton, 0)
	navButtons = append(navButtons, f.bFactory.CreateClearCartButton())
	navButtons = append(navButtons, f.bFactory.CreateOrderCreationButton())

	buttons = append(buttons, navButtons)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = buttons

	return inlineKeyboard
}

func (f *MurkupFactory) CreateOrdersListMenu(orders []api.OrderShort) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()

	for _, order := range orders {
		inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, f.bFactory.CreateOrderShortButtonGroup(order))
	}

	return inlineKeyboard
}

func (f *MurkupFactory) CreateOrderInfo(order api.Order) tgbotapi.InlineKeyboardMarkup {
	buttons := f.bFactory.CreateOrderButtonGroup(order.Id)

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
