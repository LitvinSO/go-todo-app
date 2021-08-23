package main

import (
	"fmt"
	"os"

	"github.com/LitvinSO/go-todo-app"
	"github.com/LitvinSO/go-todo-app/pkg/handler"
	"github.com/LitvinSO/go-todo-app/pkg/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	fmt.Println("get config - ", viper.GetString("db.host"))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loaging env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		TimeZone: viper.GetString("db.timezone"),
	})

	if err != nil {
		logrus.Fatalf("error initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)

	if repos == nil {
		logrus.Fatalf("error repos ")
	}

	handlers := new(handler.Handler)
	srv := new(todo.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRouts())
	if err != nil {
		logrus.Fatalf("error occured while running http Server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
