package dependencies

import (
	"hello-fuego/cmd/api/core/usecases"
	"hello-fuego/cmd/api/entrypoints"
	"hello-fuego/cmd/api/entrypoints/handlers"
)

type HandlerContainer struct {
	HelloWorld entrypoints.Handler
}

func Start() *HandlerContainer {
	helloWorldUseCase := &usecases.HelloWorldUseCaseImpl{}

	apiHandlers := HandlerContainer{}

	apiHandlers.HelloWorld = &handlers.HelloWorldHandler{
		UseCase: helloWorldUseCase,
	}

	return &apiHandlers
}
