package factories

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	back              = "back"
	Back_CurrentId    = "currentId"
	ProductSelect_Id  = "productId"
	CategorySelect_Id = "categoryId"
)

type ButtonsFactory interface {
	CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton
	CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton
	CreateBackButton(currentId int) tgbotapi.InlineKeyboardButton
}

type MainButtonsFactory struct {
}

func NewMainButtonsFactory() *MainButtonsFactory {
	return &MainButtonsFactory{}
}

func (f *MainButtonsFactory) CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(categoryName, c.CategoryPrefix+"_"+formatData(CategorySelect_Id, strconv.Itoa(categoryId)))
}
func (f *MainButtonsFactory) CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(productName, c.ProductPrefix+"_"+formatData(ProductSelect_Id, strconv.Itoa(productId)))
}
func (f *MainButtonsFactory) CreateBackButton(currentId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(back, c.BackPrefix+"_"+formatData(Back_CurrentId, strconv.Itoa(currentId))+"|"+formatData("status", "ok"))
}

func formatData(key string, value string) string {
	return key + "=" + value
}

type Button struct {
	Type c.CallBackType
	Data map[string]string
}
