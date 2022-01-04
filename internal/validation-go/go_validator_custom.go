package validation_go

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"idaman.id/storage/internal/config"
)

type CustomValidator = func(fl validator.FieldLevel) bool
type CustomTagName = func(field reflect.StructField) string

func NewTagNameFunc() CustomTagName {
	return func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	}
}

func NewValidFileSizeRule(configGetter config.Getter) CustomValidator {
	return func(fl validator.FieldLevel) bool {
		size := fl.Field().Interface().(int64)
		fileSize := int(size)

		minSize := configGetter.GetInt("MIN_FILE_SIZE")
		maxSize := configGetter.GetInt("MAX_FILE_SIZE")
		isSizeValid := fileSize >= minSize && fileSize <= maxSize

		return isSizeValid
	}
}

// func NewValidProviderRule() CustomValidator {
// 	return func(fl validator.FieldLevel) bool {
// 		value := fl.Field().Interface().(string)

// 		isProviderValid := value == "local"
// 		return isProviderValid
// 	}
// }

// func NewValidFileAmountRule(configGetter config.Getter) CustomValidator {
// 	return func(fl validator.FieldLevel) bool {

// 		var totalFile int
// 		value := fl.Field().Interface()
// 		switch reflect.TypeOf(value).Kind() {
// 		case reflect.Slice:
// 			totalFile = reflect.ValueOf(value).Len()
// 		}

// 		minAmount := configGetter.GetInt("MIN_UPLOADED_FILE")
// 		maxAmount := configGetter.GetInt("MAX_UPLOADED_FILE")
// 		isAmountValid := totalFile >= minAmount && totalFile <= maxAmount

// 		return isAmountValid
// 	}
// }
