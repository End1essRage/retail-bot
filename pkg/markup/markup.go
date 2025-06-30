package markup

import (
	c "github.com/end1essrage/retail-bot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Back_CurrentId     = "curid"
	Back_IsProduct     = "ip"
	Product_Id         = "pid"
	Product_Name       = "pn"
	Category_Id        = "catid"
	Order_Id           = "oid"
	Order_TargetStatus = "ts"
	Order_ClientChatId = "cid"
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
