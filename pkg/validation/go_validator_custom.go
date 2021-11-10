package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"idaman.id/storage/pkg/config"
)

var tagNameFunc = func(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

var validProviderRule = func(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(string)

	isProviderValid := value == "local"
	return isProviderValid
}

var validFileAmountRule = func(fl validator.FieldLevel) bool {

	var totalFile int
	value := fl.Field().Interface()
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		totalFile = reflect.ValueOf(value).Len()
	}

	minAmount := config.Service.GetInt("MIN_UPLOADED_FILE")
	maxAmount := config.Service.GetInt("MAX_UPLOADED_FILE")
	isAmountValid := totalFile >= minAmount && totalFile <= maxAmount

	return isAmountValid
}

var validFileSizeRule = func(fl validator.FieldLevel) bool {
	size := fl.Field().Interface().(uint64)
	fileSize := int(size)

	minSize := config.Service.GetInt("MIN_FILE_SIZE")
	maxSize := config.Service.GetInt("MAX_FILE_SIZE")
	isSizeValid := fileSize >= minSize && fileSize <= maxSize

	return isSizeValid
}
