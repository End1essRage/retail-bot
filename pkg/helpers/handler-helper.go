package helpers

import (
	"fmt"
	"strings"

	c "github.com/end1essrage/retail-bot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handling buttons
func GetCallBackTypeAndData(callback *tgbotapi.CallbackQuery) (*CallbackData, error) {
	cbType := strings.Split(callback.Data, "_")[0]
	result := &CallbackData{}
	switch cbType {
	case c.CategoryPrefix:
		result.Type = c.CategorySelect
	case c.ProductPrefix:
		result.Type = c.ProductSelect
	case c.BackPrefix:
		result.Type = c.Back
	default:
		return nil, fmt.Errorf("unknown command")
	}

	result.Data = formatData(strings.Split(callback.Data, "_")[1])
	return result, nil
}

func formatData(data string) map[string]string {
	result := make(map[string]string)

	items := strings.Split(data, "|")
	for _, i := range items {
		key := strings.Split(i, "=")[0]
		value := strings.Split(i, "=")[1]
		result[key] = value
	}

	return result
}

type CallbackData struct {
	Type c.CallBackType
	Data map[string]string
}
