package usecases

import "context"

type ResponseTest struct {
	Hello string `json:"hello"`
}

type HelloWorldUseCase interface {
	Execute(ctx context.Context) (ResponseTest, error)
}

func NewHelloWorldUseCase() HelloWorldUseCaseImpl {
	return HelloWorldUseCaseImpl{}
}

type HelloWorldUseCaseImpl struct {
}

func (imp *HelloWorldUseCaseImpl) Execute(ctx context.Context) (ResponseTest, error) {
	return ResponseTest{Hello: "world"}, nil
}
