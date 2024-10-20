package markup

import (
	c "github.com/end1essrage/retail-bot/pkg"
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

func formatData(key string, value string) string {
	return key + c.FlagSeparator + value
}
