package app

import (
	"github.com/gin-gonic/gin"
	"hello-fuego/cmd/api/infrastructure/dependencies"
)

func ConfigureMappings(router *gin.Engine, apiHandlers *dependencies.HandlerContainer) *gin.Engine {
	router.GET("/hello-world", apiHandlers.HelloWorld.Handle)

	return router
}
