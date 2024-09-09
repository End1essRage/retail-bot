package main

import (
	"flag"
	"os"

	c "github.com/end1essrage/retail-bot/pkg"
	"github.com/end1essrage/retail-bot/pkg/api"
	"github.com/end1essrage/retail-bot/pkg/factories"
	"github.com/end1essrage/retail-bot/pkg/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Token string
	Env   string
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	Env = os.Getenv("ENVIRONMENT")

	if Env == "" {
		if err := godotenv.Load(); err != nil {
			logrus.Fatalf("error while reading environment %s", err.Error())
		}
	}

	Env = os.Getenv("ENVIRONMENT")

	if Env == "" {
		logrus.Fatal("cant set environment")
	}

	logrus.Info("ENVIRONMENT IS " + Env)

	setToken()
	//Config Handling

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config %s", err.Error())
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		logrus.Panic(err)
	}

	bot.Debug = true

	api := api.NewApi(viper.GetString("api_host"))
	bFactory := factories.NewMainButtonsFactory()
	mFactory := factories.NewMurkupFactory(bFactory)
	handler := handler.NewTgHandler(bot, api, bFactory, mFactory)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handler.Handle(&update)
	}
}

func initConfig() error {
	if Env == c.ENV_LOCAL {
		viper.SetConfigName("config_local")
	}
	if Env == c.ENV_DEV {
		viper.SetConfigName("config_pod")
	}

	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")

	return viper.ReadInConfig()
}

func setToken() {
	if Env == c.ENV_LOCAL {
		flag.StringVar(&Token, "t", "", "Bot Token")
		flag.Parse()
	}

	if Env == c.ENV_DEV {
		Token = os.Getenv("TOKEN")
	}
}
