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

func GetRole(userName string) c.UserRole {
	aNames := viper.GetString("manager_names")
	names := strings.Split(aNames, " ")
	logrus.Warn(aNames)
	role := c.Client
	for _, n := range names {
		logrus.Info(n)
		if n == userName {
			role = c.Manager
			logrus.Warn("manager")
			break
		}
	}

	return role
}

// handling buttons
func GetCallBackTypeAndData(callback *tgbotapi.CallbackQuery) *CallbackData {
	cbType := strings.Split(callback.Data, c.TypeSeparator)[0]
	result := &CallbackData{}
	result.Type = c.CallBackType(cbType)
	if len(strings.Split(callback.Data, c.TypeSeparator)) > 1 {
		var data = strings.Split(callback.Data, c.TypeSeparator)[1]
		if data != "" {
			result.Data = formatData(data)
		}
	}
	return result
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
