package main

import (
	"flag"
	"log"
	"os"

	"github.com/ThailanTec/poc-serasa/adapter/handler"
	"github.com/ThailanTec/poc-serasa/adapter/repository/postgres"
	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	"github.com/ThailanTec/poc-serasa/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	repo = flag.String("db", "postgres", "Database for storing messages")
	// redisHost   = "localhost:6379"
	httpHandler *handler.HTTPHandler
	svc         *services.UserService
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//migrate := flag.Bool("disable-migrate", false, "Migrate internal databases automatically")

	metadata := postgres.NewPG(domain.DBSettings{
		Host:               os.Getenv("DATA_DB_HOST"),
		User:               os.Getenv("DATA_DB_USER"),
		Password:           os.Getenv("DATA_DB_PASS"),
		DBName:             os.Getenv("DATA_DB_NAME"),
		Port:               os.Getenv("DATA_DB_PORT"),
		SSL:                os.Getenv("DATA_DB_SSL"),
		TimeZone:           os.Getenv("DATA_DB_TIMEZONE"),
		DisableAutoMigrate: false,
	})

	svc = services.NewServices(metadata)

	InitRoutes()
}

func InitRoutes() {
	router := gin.Default()
	handler := handler.NewHttpHandler(*svc)
	router.GET("/getUByID/:id", handler.GetUserByID)
	router.DELETE("/delete/:id", handler.DeleteUserByID)
	router.POST("/create", handler.CreateUser)

	router.Run(":9000")
}
