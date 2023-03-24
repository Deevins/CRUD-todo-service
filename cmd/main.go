package main

import (
	"context"
	todoserver "github.com/deevins/todo-restAPI"
	"github.com/deevins/todo-restAPI/internal/handler"
	repository "github.com/deevins/todo-restAPI/internal/repository"
	"github.com/deevins/todo-restAPI/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// TODO: change logrus lib to uber logger(zap)
// https://github.com/uber-go/zap
// TODO: add unique id generation for TODO-ITEM and TODO-LIST

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	if err := initConfig(); err != nil {
		logrus.Fatalf("can not read config file %s", err.Error())
	}

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	port := viper.GetString("port")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize DB, %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoserver.Server)
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server %s", err.Error())
		}
	}()
	logrus.Println("App started successfully")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Println("Application is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error on server while shutting down: [%s]", err)
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error on Database closing connection: [%s]", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
