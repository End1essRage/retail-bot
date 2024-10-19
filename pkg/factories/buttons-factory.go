package factories

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Back_CurrentId = "currentId"
	Back_IsProduct = "isProduct"
	Product_Id     = "productId"
	Product_Name   = "productName"
	Category_Id    = "categoryId"
	Order_Id       = "orderId"
)

// разделитть на более мелкие
type MenuButtonsFactory interface {
	// Кнопка категории для проваливания в нее
	CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton

	// Кнопка товара для проваливания в его карточку
	CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton

	// Кнопка назад для навигации по меню
	CreateBackButton(currentId int, isProduct bool) tgbotapi.InlineKeyboardButton

	// Кнопка добавления товара в корзину
	CreateAddButton(productId int, productName string) tgbotapi.InlineKeyboardButton
}

type CartButtonsFactory interface {
	//кнопка инкремента и дкермента кол-ва товара и кнопка с названием товара
	CreatePositionButtonGroup(productId int, productName string, amount int) [][]tgbotapi.InlineKeyboardButton

	//кнопка создания заказа
	CreateOrderCreationButton() tgbotapi.InlineKeyboardButton

	//кнопка очистки корзины
	CreateClearCartButton() tgbotapi.InlineKeyboardButton
}

type OrderButtonsFactory interface {
	CreateOrderShortButtonGroup(order api.OrderShort) []tgbotapi.InlineKeyboardButton
	CreateOrderButtonGroup(orderId int) []tgbotapi.InlineKeyboardButton
}

type MainButtonsFactory struct {
}

func NewMainButtonsFactory() *MainButtonsFactory {
	return &MainButtonsFactory{}
}

func (f *MainButtonsFactory) CreateCategorySelectButton(categoryName string, categoryId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(categoryName,
		string(c.CategorySelect)+c.TypeSeparator+formatData(Category_Id, strconv.Itoa(categoryId)))
}

func (f *MainButtonsFactory) CreateAddButton(productId int, productName string) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("add",
		string(c.ProductAdd)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId))+c.DataSeparator+formatData(Product_Name, productName))
}

func (f *MainButtonsFactory) CreateProductSelectButton(productName string, productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(productName,
		string(c.ProductSelect)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateBackButton(parentId int, isProduct bool) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("back to "+strconv.Itoa(parentId),
		string(c.Back)+c.TypeSeparator+formatData(Back_CurrentId, strconv.Itoa(parentId))+c.DataSeparator+formatData(Back_IsProduct, strconv.FormatBool(isProduct)))
}

func (f *MainButtonsFactory) CreateIncrementPositionButton(productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("(+)",
		string(c.ProductIncrement)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateDecrementPositionButton(productId int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("(-)",
		string(c.ProductDecrement)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateNamePositionButton(productId int, productName string) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(productName,
		string(c.ProductName)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))
}

func (f *MainButtonsFactory) CreateAmountPositionButton(productId int, amount int) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(amount),
		string(c.ProductAmount)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))
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

func (f *MainButtonsFactory) CreateOrderShortButtonGroup(order api.OrderShort) []tgbotapi.InlineKeyboardButton {
	/*
		0 - new - cancel
		1 - accepted - cancel
		2 - completed - rate
		3 - cancelled - repeate / none
	*/

	// дата состав статус
	result := make([]tgbotapi.InlineKeyboardButton, 0)
	/*
		timeButton := tgbotapi.NewInlineKeyboardButtonData(date.String(), "nnn"+c.TypeSeparator)d
		itemsButton := tgbotapi.NewInlineKeyboardButtonData(name, "nnn"+c.TypeSeparator)
		statusButton := tgbotapi.NewInlineKeyboardButtonData(status, "nnn"+c.TypeSeparator)

		result = append(result, timeButton)
		result = append(result, itemsButton)
		result = append(result, statusButton)
		.Format("2006-01-02 15:04:05")
	*/
	orderButton := tgbotapi.NewInlineKeyboardButtonData(order.DateCreation.Format("02-01 15:04")+" | "+"short items"+" | "+order.StatusName, string(c.OrderShortOpen)+c.TypeSeparator+
		formatData(Order_Id, strconv.Itoa(order.Id)))

	result = append(result, orderButton)

	return result
}

// back and cancel
func (f *MainButtonsFactory) CreateOrderButtonGroup(orderId int) []tgbotapi.InlineKeyboardButton {
	result := make([]tgbotapi.InlineKeyboardButton, 0)
	//от статуса будет зависеть наличие кнопки cancel
	backButton := tgbotapi.NewInlineKeyboardButtonData("back", string(c.OrderBackToList)+c.TypeSeparator)
	cancelButton := tgbotapi.NewInlineKeyboardButtonData("cancel", string(c.OrderCancel)+c.TypeSeparator+
		formatData(Order_Id, strconv.Itoa(orderId)))

	result = append(result, backButton)
	result = append(result, cancelButton)

	return result
}

func (f *MainButtonsFactory) CreateOrderCreationButton() tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("create",
		string(c.CreateOrder)+c.TypeSeparator)
}

func (f *MainButtonsFactory) CreateClearCartButton() tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData("clear",
		string(c.ClearCart)+c.TypeSeparator)
}

func formatData(key string, value string) string {
	return key + c.FlagSeparator + value
}
