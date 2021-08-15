package main

import (
	"log"

	"github.com/LitvinSO/go-todo-app"
	"github.com/LitvinSO/go-todo-app/pkg/handler"
	"github.com/LitvinSO/go-todo-app/pkg/repository"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "Paz3585@Paz3585",
		DBname:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("error initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)

	if repos == nil {
		log.Fatalf("error repos ")
	}

	handlers := new(handler.Handler)
	srv := new(todo.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRouts())
	if err != nil {
		log.Fatalf("error occured while running http Server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
