package validators_test

import (
	"github.com/gin-gonic/gin"
	"hello-fuego/cmd/api/config/validators"
	"testing"
)

func TestBindAndValidate(t *testing.T) {
	type args struct {
		c   *gin.Context
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validators.BindAndValidate(tt.args.c, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("BindAndValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
