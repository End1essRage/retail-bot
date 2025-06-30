package markup

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type MenuMarkup interface {
	MenuRoot(categories []api.Category) tgbotapi.InlineKeyboardMarkup
	MenuCategoriesList(categories []api.Category) tgbotapi.InlineKeyboardMarkup
	MenuProductsList(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup
	MenuProductForm(Product api.Product) tgbotapi.InlineKeyboardMarkup
}

// client
type ClientMarkup struct {
}

func NewClientMarkup() *ClientMarkup {
	return &ClientMarkup{}
}

func (m *ClientMarkup) MenuRoot(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		bt := createCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (m *ClientMarkup) MenuCategoriesList(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
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

func (m *ClientMarkup) MenuProductsList(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range Products {
		btn := tgbotapi.NewInlineKeyboardButtonData(p.Name,
			string(c.ProductSelect)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(p.Id)))

		buttons = append(buttons, btn)
	}

	buttons = append(buttons, createBackButton(categoryId, false))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (m *ClientMarkup) MenuProductForm(Product api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	logrus.Warning(Product)

	addButton := tgbotapi.NewInlineKeyboardButtonData("add",
		string(c.ProductAdd)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(Product.Id)))

	buttons = append(buttons, addButton)
	buttons = append(buttons, createBackButton(Product.CategoryId, true))

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

// manager
type ManagerMarkup struct {
}

func NewManagerMarkup() *ManagerMarkup {
	return &ManagerMarkup{}
}

func (m *ManagerMarkup) MenuRoot(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, c := range categories {
		bt := createCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
	}

	buttons = append(buttons, createIAMADMINButton())

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (m *ManagerMarkup) MenuCategoriesList(categories []api.Category) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	parentId := 0

	for _, c := range categories {
		bt := createCategorySelectButton(c.Name, c.Id)
		buttons = append(buttons, bt)
		parentId = c.Parent
	}

	buttons = append(buttons, createBackButton(parentId, false))
	buttons = append(buttons, createIAMADMINButton())

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (m *ManagerMarkup) MenuProductsList(categoryId int, Products []api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	for _, p := range Products {
		btn := tgbotapi.NewInlineKeyboardButtonData(p.Name,
			string(c.ProductSelect)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(p.Id)))

		buttons = append(buttons, btn)
	}

	buttons = append(buttons, createBackButton(categoryId, false))

	buttons = append(buttons, createIAMADMINButton())

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}

func (m *ManagerMarkup) MenuProductForm(Product api.Product) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	logrus.Warning(Product)

	addButton := tgbotapi.NewInlineKeyboardButtonData("add",
		string(c.ProductAdd)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(Product.Id)))

	buttons = append(buttons, addButton)
	buttons = append(buttons, createBackButton(Product.CategoryId, true))

	buttons = append(buttons, createIAMADMINButton())

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = groupButtons(buttons, 1)

	return inlineKeyboard
}
func createCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(categoryName,
		string(c.CategorySelect)+c.TypeSeparator+formatData(Category_Id, strconv.Itoa(categoryId)))
}

func createBackButton(parentId int, isProduct bool) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("back",
		string(c.Back)+c.TypeSeparator+formatData(Back_CurrentId, strconv.Itoa(parentId))+c.DataSeparator+formatData(Back_IsProduct, strconv.FormatBool(isProduct)))
}

func createIAMADMINButton() tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("IAMADMIN", "appapap_sss=1")
}
