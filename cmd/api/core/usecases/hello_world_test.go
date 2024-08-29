package usecases_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hello-fuego/cmd/api/core/usecases"
	"testing"
)

func TestHelloWorldUseCaseImpl_Execute(t *testing.T) {
	tests := []struct {
		name string
		want usecases.ResponseTest
	}{
		{
			name: "Test HelloWorldUseCaseImpl_Execute",
			want: usecases.ResponseTest{Hello: "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			imp := &usecases.HelloWorldUseCaseImpl{}

			got, _ := imp.Execute(ctx)

			assert.Equalf(t, tt.want, got, "Execute() = %v, want %v", got, tt.want)
		})
	}
}
