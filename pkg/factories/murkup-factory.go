package factories

import (
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MenuMurkupFactory interface {
	//изначальный список категорий
	CreateRootMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup
	//список категорий
	CreateCategorySelectMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup
	//список товаров
	CreateProductSelectMenu(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup
	// карточка товара
	CreateProductMenu(Product api.Product) tgbotapi.InlineKeyboardMarkup
}

type CartMurkupFactory interface {
	//меню корзины
	CreateCartMenu(positions []service.Position) tgbotapi.InlineKeyboardMarkup
}

type UserMurkupFactory struct {
	mFactory MenuButtonsFactory
	cFactory CartButtonsFactory
}

type AdminMurkupFactory struct {
	mFactory MenuButtonsFactory
}

func NewUserMurkupFactory() *UserMurkupFactory {
	bfactory := NewMainButtonsFactory()
	return &UserMurkupFactory{mFactory: bfactory, cFactory: bfactory}
}

func (f *UserMurkupFactory) CreateRootMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		bt := f.mFactory.CreateCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *UserMurkupFactory) CreateCategorySelectMenu(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	parentId := 0

	for _, c := range categories {
		bt := f.mFactory.CreateCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
		parentId = c.Parent
	}

	buttons = append(buttons, f.mFactory.CreateBackButton(parentId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *UserMurkupFactory) CreateProductSelectMenu(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range Products {
		buttons = append(buttons, f.mFactory.CreateProductSelectButton(p.Name, p.Id))
	}

	buttons = append(buttons, f.mFactory.CreateBackButton(categoryId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *UserMurkupFactory) CreateProductMenu(Product api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, f.mFactory.CreateAddButton(Product.Id, Product.Name))
	buttons = append(buttons, f.mFactory.CreateBackButton(Product.CategoryId, true))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (f *UserMurkupFactory) CreateCartMenu(positions []service.Position) tgbotapi.InlineKeyboardMarkup {
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)

	for _, pos := range positions {
		positionButtons := f.cFactory.CreatePositionButtonGroup(pos.Product.Id, pos.Product.Name, pos.Count)
		buttons = append(buttons, positionButtons[0])
		buttons = append(buttons, positionButtons[1])
	}

	navButtons := make([]tgbotapi.InlineKeyboardButton, 0)
	navButtons = append(navButtons, f.cFactory.CreateClearCartButton())
	navButtons = append(navButtons, f.cFactory.CreateOrderCreationButton())

	buttons = append(buttons, navButtons)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = buttons

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
