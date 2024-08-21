package main

import (
	"github.com/gin-gonic/gin"
	"hello-fuego/cmd/api/app"
	"hello-fuego/cmd/api/infrastructure/dependencies"
	"log"
	"os"
)

type Teste struct {
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	handlers := dependencies.Start()

	app.ConfigureMappings(router, handlers)

	err := router.Run(":" + port)

	if err != nil {
		logger := log.Logger{}

		logger.Println(err.Error())
	}
}
