package helpers

import (
	"fmt"
	"strings"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
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
	case c.ProductAddPrefix:
		result.Type = c.ProductAdd
	case c.BackPrefix:
		result.Type = c.Back
	case c.ProductIncrementPrefix:
		result.Type = c.ProductIncrement
	case c.ProductDecrementPrefix:
		result.Type = c.ProductDecrement
	case c.ProductNamePrefix:
		result.Type = c.ProductName
	case c.ProductAmountPrefix:
		result.Type = c.ProductAmount
	case c.ClearCartPrefix:
		result.Type = c.ClearCart
	case c.CreateOrderPrefix:
		result.Type = c.CreateOrder
	default:
		return nil, fmt.Errorf("unknown command")
	}

	if len(strings.Split(callback.Data, "_")) > 1 {
		result.Data = formatData(strings.Split(callback.Data, "_")[1])
	}

	return result, nil
}

func FilterRootCategories(categories []api.Category) []api.Category {
	categoriesFiltered := make([]api.Category, 0)
	for _, cat := range categories {
		if cat.Parent == 0 {
			categoriesFiltered = append(categoriesFiltered, cat)
		}
	}

	return categoriesFiltered
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
