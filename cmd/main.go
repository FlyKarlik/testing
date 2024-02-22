package main

import (
	"log"
	"os"
	"test_task/repository"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	orderIDs := os.Args[1:]
	if len(orderIDs) == 0 {
		log.Fatal("there are is not enough arguments in command line")
	}
	if err := InitConfig(); err != nil {
		log.Fatalf("error for initilizing config file: %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initilize db: %s", err.Error())
	}
	store := repository.NewDb(db)
	if err := store.GetOrders(orderIDs); err != nil {
		log.Fatal(err)
	}
}
func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
