package main

import (
	"Telegram-Market/api/pkg/handler"
	"Telegram-Market/api/pkg/repository"
	"Telegram-Market/api/pkg/service"
	"Telegram-Market/api/serverAPI"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initial config: %v", err.Error())
	}

	db, err := repository.NewSqliteDB()

	storageDB := repository.NewStorage(db, writer)
	services := service.NewService(storageDB)
	handlerAPI := handler.NewHandler(services, writer)

	srv := new(serverAPI.Server)

	err = srv.Run(viper.GetString("port"), handlerAPI.InitRoutes())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
