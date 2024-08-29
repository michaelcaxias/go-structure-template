package handlers

import (
	"github.com/gin-gonic/gin"
	"hello-fuego/cmd/api/config/validators"
	"hello-fuego/cmd/api/core/contracts"
	"hello-fuego/cmd/api/core/usecases"
	"hello-fuego/cmd/api/errors"
	"net/http"
)

type HelloWorldHandler struct {
	UseCase usecases.HelloWorldUseCase
}

func (handler *HelloWorldHandler) Handle(c *gin.Context) {
	err := handler.handle(c)

	if err != nil {
		c.JSON(err.Status, err)
	}
}

func (handler *HelloWorldHandler) handle(c *gin.Context) *errors.APIError {
	ctx := c.Request.Context()

	req := contracts.HelloWorldRequest{}

	causes, err := validators.BindAndValidate(c, &req)

	if err != nil {
		return errors.NewBadRequestError("some fields are invalid", causes)
	}

	r, err := handler.UseCase.Execute(ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, r)

	return nil
}
