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
	Product_Name   = "productName"
	Category_Id    = "categoryId"
)

//разделитть на более мелкие

type ButtonsFactory interface {
	CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton
	CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton
	CreateBackButton(currentId int, isProduct bool) tgbotapi.InlineKeyboardButton
	CreateAddButton(productId int, productName string) tgbotapi.InlineKeyboardButton
	//CreateIncrementPositionButton(productId int) tgbotapi.InlineKeyboardButton
	//CreateDecrementPositionButton(productId int) tgbotapi.InlineKeyboardButton
	//CreateNamePositionButton(productId int, productName string) tgbotapi.InlineKeyboardButton
	//CreateAmountPositionButton(productId int, amount int) tgbotapi.InlineKeyboardButton
	CreatePositionButtonGroup(productId int, productName string, amount int) [][]tgbotapi.InlineKeyboardButton
	CreateOrderCreationButton() tgbotapi.InlineKeyboardButton
	CreateClearCartButton() tgbotapi.InlineKeyboardButton
}

type MainButtonsFactory struct {
}

func NewMainButtonsFactory() *MainButtonsFactory {
	return &MainButtonsFactory{}
}

func (f *MainButtonsFactory) CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(categoryName, c.CategoryPrefix+"_"+formatData(Category_Id, strconv.Itoa(categoryId)))
}

func (f *MainButtonsFactory) CreateAddButton(productId int, productName string) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("add", c.ProductAddPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId))+"|"+formatData(Product_Name, productName))
}

func (f *MainButtonsFactory) CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(productName, c.ProductPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateBackButton(parentId int, isProduct bool) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("back to "+strconv.Itoa(parentId), c.BackPrefix+"_"+formatData(Back_CurrentId, strconv.Itoa(parentId))+"|"+formatData(Back_IsProduct, strconv.FormatBool(isProduct)))
}

func (f *MainButtonsFactory) CreateIncrementPositionButton(productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("(+)", c.ProductIncrementPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateDecrementPositionButton(productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("(-)", c.ProductDecrementPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateNamePositionButton(productId int, productName string) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(productName, c.ProductNamePrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateAmountPositionButton(productId int, amount int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(amount), c.ProductAmountPrefix+"_"+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreatePositionButtonGroup(productId int, productName string, amount int) [][]tgbotapi.InlineKeyboardButton {
	result := make([][]tgbotapi.InlineKeyboardButton, 0)

	resultNameRow := make([]tgbotapi.InlineKeyboardButton, 0)
	resultNameRow = append(resultNameRow, f.CreateNamePositionButton(productId, productName))

	resultButtonsRow := make([]tgbotapi.InlineKeyboardButton, 0)
	resultButtonsRow = append(resultButtonsRow, f.CreateDecrementPositionButton(productId))
	resultButtonsRow = append(resultButtonsRow, f.CreateAmountPositionButton(productId, amount))
	resultButtonsRow = append(resultButtonsRow, f.CreateIncrementPositionButton(productId))

	result = append(result, resultNameRow)
	result = append(result, resultButtonsRow)
	return result
}

func (f *MainButtonsFactory) CreateOrderCreationButton() tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("create", c.CreateOrderPrefix+"_")
}

func (f *MainButtonsFactory) CreateClearCartButton() tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("clear", c.ClearCartPrefix+"_")
}

func formatData(key string, value string) string {
	return key + "=" + value
}
