package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entras "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

var (
	Validator = validator.New()
	transl    ut.Translator
)

const (
	QueryParam = "form"
	Header     = "header"
	Body       = "json"
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

func BindAndValidate(c *gin.Context, obj interface{}) (causes []map[string]string, err error) {
	causesMap := make([]map[string]string, 0)

	if err := c.ShouldBind(obj); err != nil {
		return causesMap, fmt.Errorf("error binding request: %w", err)
	}

	if err := c.ShouldBindHeader(obj); err != nil {
		return causesMap, fmt.Errorf("error binding request headers: %w", err)
	}

	if err := Validator.Struct(obj); err != nil {
		allCauses := GetCauses(obj, err)

		return allCauses, fmt.Errorf("error binding request: %w", err)
	}

	return causesMap, nil
}

func GetCauses(obj interface{}, err error) []map[string]string {
	if err == nil {
		return nil
	}

	causes := make([]map[string]string, 0)
	reflected := reflect.ValueOf(obj).Elem()

	if reflected.Kind() == reflect.Ptr {
		reflected = reflected.Elem()
	}

	for _, validationErr := range err.(validator.ValidationErrors) {
		fieldName := getFieldName(reflected, validationErr)
		message := validationErr.Translate(transl)

		field, ok := reflected.Type().FieldByName(validationErr.StructField())
		errorType := "body"
		if ok {
			errorType = getType(field.Tag)
		}

		cause := map[string]string{
			"field":   fieldName,
			"message": message,
			"type":    errorType,
		}
		causes = append(causes, cause)
	}

	return causes
}

func getFieldName(reflected reflect.Value, validationErr validator.FieldError) string {
	field, _ := reflected.Type().FieldByName(validationErr.StructField())
	fieldName := field.Tag.Get("json")

	if fieldName == "" {
		fieldName = validationErr.StructField()
	}

	return fieldName
}

func getType(tag reflect.StructTag) string {
	switch {
	case tag.Get(Header) != "":
		return "header"
	case tag.Get(QueryParam) != "":
		return "query"
	default:
		return "body"
	}
}
