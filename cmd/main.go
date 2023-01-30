package main

import (
	todoserver "github.com/deevins/todo-restAPI"
	"github.com/deevins/todo-restAPI/pkg/handler"
	"github.com/deevins/todo-restAPI/pkg/repository"
	"github.com/deevins/todo-restAPI/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

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
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
