package handlers

import (
	"github.com/gin-gonic/gin"
	"hello-fuego/cmd/api/core/usecases"
	"net/http"
)

type HelloWorldHandler struct {
	UseCase usecases.HelloWorldUseCase
}

func (handler *HelloWorldHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	r, err := handler.UseCase.Execute(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, r)
}
