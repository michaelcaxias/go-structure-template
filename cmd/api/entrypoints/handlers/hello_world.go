package handlers

import (
	"github.com/gin-gonic/gin"
	"hello-fuego/cmd/api/config/validators"
	"hello-fuego/cmd/api/core/contracts"
	"hello-fuego/cmd/api/core/usecases"
	"net/http"
)

type HelloWorldHandler struct {
	UseCase usecases.HelloWorldUseCase
}

func (handler *HelloWorldHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	req := contracts.HelloWorldRequest{}

	err := validators.BindAndValidate(c, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := handler.UseCase.Execute(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, r)
}
