package factories

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	back           = "back"
	Back_CurrentId = "currentId"
	Back_IsProduct = "isproduct"
	Product_Id     = "productId"
	Category_Id    = "categoryId"
)

type ButtonsFactory interface {
	CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton
	CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton
	CreateBackButton(currentId int, isProduct bool) tgbotapi.InlineKeyboardButton
	CreateAddButton(productId int) tgbotapi.InlineKeyboardButton
}

type MainButtonsFactory struct {
}

func NewMainButtonsFactory() *MainButtonsFactory {
	return &MainButtonsFactory{}
}

func (f *MainButtonsFactory) CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(categoryName, c.CategoryPrefix+"_"+formatData(Category_Id, strconv.Itoa(categoryId)))
}

func (f *MainButtonsFactory) CreateAddButton(productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("add", c.ProductAddPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(productName, c.ProductPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateBackButton(parentId int, isProduct bool) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("back to "+strconv.Itoa(parentId), c.BackPrefix+"_"+formatData(Back_CurrentId, strconv.Itoa(parentId))+"|"+formatData(Back_IsProduct, strconv.FormatBool(isProduct)))
}

func formatData(key string, value string) string {
	return key + "=" + value
}

type Button struct {
	Type c.CallBackType
	Data map[string]string
}
