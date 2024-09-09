package factories

import (
	"github.com/end1essrage/retail-bot/pkg/api"
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

func (f *MurkupFactory) CreateCategorySelectMenu(currentId int, categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		bt := f.bFactory.CreateCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
	}

	buttons = append(buttons, f.bFactory.CreateBackButton(currentId))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *MurkupFactory) CreateProductSelectMenu(currentId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range Products {
		bt := f.bFactory.CreateProductSelectButton(p.Name, p.Id)
		buttons = append(buttons, bt)
	}

	buttons = append(buttons, f.bFactory.CreateBackButton(currentId))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

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
