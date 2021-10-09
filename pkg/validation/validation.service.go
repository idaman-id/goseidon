package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"idaman.id/storage/pkg/app"
	"idaman.id/storage/pkg/config"
	"idaman.id/storage/pkg/file"
)

func ValidateStruct(param ValidationStructDto) *ValidationError {
	translator := createTranslator(param.Locale)
	validate, validatorErr := createValidator(translator, param.Locale)
	if validatorErr != nil {
		panic("Failed to create validator")
	}

	validationResult := validate.Struct(param.Struct)
	isDataValid := validationResult == nil
	if isDataValid {
		return nil
	}

	validationError := ValidationError{
		Message: app.STATUS_INVALID_DATA,
	}

	errors := validationResult.(validator.ValidationErrors)
	for _, error := range errors {
		value := getValueAsString(error.Value())

		element := ValidationItem{
			Field:   error.Field(),
			Message: error.Translate(translator),
			Value:   value,
		}
		validationError.Items = append(validationError.Items, &element)
	}

	return &validationError
}

func ValidateRule(param ValidationRuleDto) *ValidationError {
	translator := createTranslator(param.Locale)
	validate, validatorErr := createValidator(translator, param.Locale)
	if validatorErr != nil {
		panic("Failed to create validator")
	}

	vaidationResult := validate.ValidateMap(param.Data, param.Rule)

	isDataValid := len(vaidationResult) == 0
	if isDataValid {
		return nil
	}

	validationError := ValidationError{
		Message: app.STATUS_INVALID_DATA,
	}

	for field, err := range vaidationResult {

		var errors validator.ValidationErrors
		switch err.(type) {
		case validator.ValidationErrors:
			errors = err.(validator.ValidationErrors)
		}

		for _, error := range errors {
			/*
				@note follow this github issue:
				https://github.com/go-playground/validator/issues/805
				for the moment I gonna trim the error message
				pretending that the field name is not spelled on the message
				e.g:
				- ` is a required field` -> `is a required field`
			*/
			message := strings.Trim(error.Translate(translator), " ")
			value := getValueAsString(error.Value())

			element := ValidationItem{
				Field:   field, //error.Field()
				Message: message,
				Value:   value,
			}
			validationError.Items = append(validationError.Items, &element)
		}
	}

	return &validationError
}

func getValueAsString(error interface{}) string {
	var value string
	switch error.(type) {
	case string:
		value = error.(string)
	case bool:
		bValue := error.(bool)
		value = strconv.FormatBool(bValue)
	case float32, float64:
		fValue := error.(float64)
		value = fmt.Sprintf("%f", fValue)
	case uint64, uint32, uint16, uint8:
		uintValue := error.(uint64)
		value = strconv.FormatUint(uintValue, 10)
	case int64, int32, int16, int8:
		intValue := error.(int64)
		value = strconv.FormatInt(intValue, 10)
	default:
		value = ""
	}
	return value
}

func createValidator(translator ut.Translator, locale string) (*validator.Validate, error) {

	validate := validator.New()

	registerTag(validate)

	var err error
	err = registerValidation(validate)
	if err != nil {
		return nil, err
	}

	err = registerDefaultTranslation(validate, translator, locale)
	if err != nil {
		return nil, err
	}

	err = registerTranslation(validate, translator)
	if err != nil {
		return nil, err
	}

	return validate, nil
}

func registerTag(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func registerValidation(validate *validator.Validate) error {
	var err error

	err = validate.RegisterValidation("valid_provider", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(string)

		if value == "local" {
			return true
		}
		/*
			@todo
			1. check `value` to `database.provider.unique_id`
		*/
		return false
	})
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("valid_file_type", func(fl validator.FieldLevel) bool {
		fileType := fl.Field().Interface().(string)

		result, isTypeSupported := file.SUPPORTED_TYPES[fileType]
		return result && isTypeSupported
	})
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("valid_file_amounts", func(fl validator.FieldLevel) bool {

		var length int
		value := fl.Field().Interface()
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(value)
			length = s.Len()
		}

		minLength := config.GetInt("MIN_UPLOADED_FILE")
		maxLength := config.GetInt("MAX_UPLOADED_FILE")
		isLengthValid := length >= minLength && length <= maxLength

		return isLengthValid
	})
	if err != nil {
		return err
	}

	return nil
}
