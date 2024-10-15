package helpers

import (
	"strings"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CallbackData struct {
	Type c.CallBackType
	Data map[string]string
}

func IsAdmin(userName string) bool {
	aNames := viper.GetString("admin_names")
	names := strings.Split(aNames, " ")

	var flag = false
	for _, n := range names {
		logrus.Info(n)
		if n == userName {
			flag = true
			break
		}
	}

	return flag
}

// handling buttons
func GetCallBackTypeAndData(callback *tgbotapi.CallbackQuery) (*CallbackData, error) {
	logrus.Warning("1")
	cbType := strings.Split(callback.Data, c.TypeSeparator)[0]
	logrus.Warning("2")
	result := &CallbackData{}
	result.Type = c.CallBackType(cbType)
	logrus.Warning("3")
	if len(strings.Split(callback.Data, c.TypeSeparator)) > 1 {
		var data = strings.Split(callback.Data, c.TypeSeparator)[1]
		if data != "" {
			result.Data = formatData(data)
		}
	}
	logrus.Warning("4")
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

	items := strings.Split(data, c.DataSeparator)
	for _, i := range items {
		key := strings.Split(i, c.FlagSeparator)[0]
		value := strings.Split(i, c.FlagSeparator)[1]
		result[key] = value
	}

	return result
}
