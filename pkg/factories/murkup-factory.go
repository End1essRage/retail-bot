package factories

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MurkupFactory struct {
	bFactory ButtonsFactory
}

func NewMurkupFactory(factory ButtonsFactory) *MurkupFactory {
	return &MurkupFactory{bFactory: factory}
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
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, pos := range positions {
		buttons = append(buttons, f.bFactory.CreateNamePositionButton(pos.Product.Id, pos.Product.Name))
		buttons = append(buttons, f.bFactory.CreateAmountPositionButton(pos.Product.Id, pos.Count))
		buttons = append(buttons, f.bFactory.CreateIncrementPositionButton(pos.Product.Id))
		buttons = append(buttons, f.bFactory.CreateDecrementPositionButton(pos.Product.Id))
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 4)

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
