package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinialeixo/backend-test/src/config/logger"
	"github.com/vinialeixo/backend-test/src/controller/routes"
)

func main() {
	logger.Info("Initializing API")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("fatal error while loading config")
	}

	router := gin.Default()

	routes.InitRouters(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
