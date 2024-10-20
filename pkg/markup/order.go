package markup

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func CreateOrdersListMenu(orders []api.OrderShort) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	logrus.Warn(orders)
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

func createOrderShortButtonGroup(order api.OrderShort) []tgbotapi.InlineKeyboardButton {
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

func CreateOrderManagerButtonGroup(orderId int) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	//от статуса будет зависеть наличие кнопки cancel
	acceptButton := tgbotapi.NewInlineKeyboardButtonData("Accept", string(c.OrderAccept)+c.TypeSeparator+
		formatData(Order_Id, strconv.Itoa(orderId)))

	cancelButton := tgbotapi.NewInlineKeyboardButtonData("Cancel", string(c.OrderCancel)+c.TypeSeparator+
		formatData(Order_Id, strconv.Itoa(orderId)))

	buttons = append(buttons, acceptButton)
	buttons = append(buttons, cancelButton)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = append(inlineKeyboard.InlineKeyboard, buttons)

	return inlineKeyboard
}

// back and cancel
func createOrderButtonGroup(orderId int) []tgbotapi.InlineKeyboardButton {
	result := make([]tgbotapi.InlineKeyboardButton, 0)
	//от статуса будет зависеть наличие кнопки cancel
	backButton := tgbotapi.NewInlineKeyboardButtonData("back", string(c.OrderBackToList)+c.TypeSeparator)
	cancelButton := tgbotapi.NewInlineKeyboardButtonData("cancel", string(c.OrderCancel)+c.TypeSeparator+
		formatData(Order_Id, strconv.Itoa(orderId)))

	result = append(result, backButton)
	result = append(result, cancelButton)

	return result
}
