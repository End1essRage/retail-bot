package markup

import (
	"strconv"

	c "github.com/end1essrage/retail-bot/pkg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateCartMenu(positions []c.Position) tgbotapi.InlineKeyboardMarkup {
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)

	for _, pos := range positions {
		positionButtons := createPositionButtonGroup(pos.Product.Id, pos.Product.Name, pos.Count)
		buttons = append(buttons, positionButtons[0])
		buttons = append(buttons, positionButtons[1])
	}

	navButtons := make([]tgbotapi.InlineKeyboardButton, 0)

	clearCartButton := tgbotapi.NewInlineKeyboardButtonData("clear",
		string(c.ClearCart)+c.TypeSeparator)

	createOrderButton := tgbotapi.NewInlineKeyboardButtonData("create",
		string(c.CreateOrder)+c.TypeSeparator)

	navButtons = append(navButtons, clearCartButton)
	navButtons = append(navButtons, createOrderButton)

	buttons = append(buttons, navButtons)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup()
	inlineKeyboard.InlineKeyboard = buttons

	return inlineKeyboard
}

func createPositionButtonGroup(productId int, productName string, amount int) [][]tgbotapi.InlineKeyboardButton {
	result := make([][]tgbotapi.InlineKeyboardButton, 0)

	nameButton := tgbotapi.NewInlineKeyboardButtonData(productName,
		string(c.ProductName)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))

	decrementButton := tgbotapi.NewInlineKeyboardButtonData("(-)",
		string(c.ProductDecrement)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))

	amountButton := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(amount),
		string(c.ProductAmount)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))

	incrementButton := tgbotapi.NewInlineKeyboardButtonData("(+)",
		string(c.ProductIncrement)+c.TypeSeparator+formatData(Product_Id, strconv.Itoa(productId)))

	resultNameRow := make([]tgbotapi.InlineKeyboardButton, 0)
	resultNameRow = append(resultNameRow, nameButton)

	resultButtonsRow := make([]tgbotapi.InlineKeyboardButton, 0)
	resultButtonsRow = append(resultButtonsRow, decrementButton)
	resultButtonsRow = append(resultButtonsRow, amountButton)
	resultButtonsRow = append(resultButtonsRow, incrementButton)

	result = append(result, resultNameRow)
	result = append(result, resultButtonsRow)
	return result
}
