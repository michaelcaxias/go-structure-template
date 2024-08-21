package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entras "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validator = validator.New()
	transl    ut.Translator
)

func init() {
	english := en.New()
	uni := ut.New(english, english)

	transl, _ = uni.GetTranslator("en")

	err := entras.RegisterDefaultTranslations(Validator, transl)

	if err != nil {
		return
	}
}

func BindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBind(obj); err != nil {
		return fmt.Errorf("error binding request: %w", err)
	}

	if err := c.ShouldBindHeader(obj); err != nil {
		return fmt.Errorf("error binding request headers: %w", err)
	}

	if err := Validator.Struct(obj); err != nil {
		// TODO: better validation here
		return fmt.Errorf("error binding request: %w", err)
	}

	return nil
}
