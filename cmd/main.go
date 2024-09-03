package main

import (
	"flag"
	"os"

	c "github.com/end1essrage/retail-bot/pkg"
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
	logrus.SetFormatter(&logrus.JSONFormatter{})

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

	logrus.Printf("Authorized on account %s", bot.Self.UserName)

	handler := handler.NewTgHandler(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handler.Handle(&update)
	}
}

func initConfig() error {
	v := viper.New()

	if Env == c.ENV_LOCAL {
		v.SetConfigName("config_local")
	}
	if Env == c.ENV_DEV {
		v.SetConfigName("config_pod")
	}

	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	return v.ReadInConfig()
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
