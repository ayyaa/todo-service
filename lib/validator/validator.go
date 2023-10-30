package validator

import (
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	alphaNumeric                    = regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
	UnmappedrrMsg map[string]string = map[string]string{
		"required":     "Value is required",
		"max":          "Value is more than max character",
		"min":          "value is less thank min character",
		"alphanumeric": "value is only alphanumeric",
		"extention":    "please check you extention file only .pdf and .txt",
	}
)

var _validator *validator.Validate

func Validate() *validator.Validate {
	return _validator
}

func InitValidator(extendedFuncs ...func(*validator.Validate)) {
	_validator = validator.New()

	// Registering the new rule "alphanumeric" ...
	err := _validator.RegisterValidation("alphanumeric", CustomValidationAlphanumeric)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	// Registering the new rule "alphanumeric" ...
	err = _validator.RegisterValidation("extention", CustomValidationExtention)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	// Registering the new rule "alphanumeric" ...
	err = _validator.RegisterValidation("enum_filter", CustomValidationEnumFilter)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	for _, extFunc := range extendedFuncs {
		extFunc(_validator)
	}
}

func ValidateErrToMapString(valErr validator.ValidationErrors) map[string]string {
	result := make(map[string]string, len(valErr))
	for _, err := range valErr {
		result[err.Field()] = fmt.Sprintf("validation: %v", err.ActualTag())
	}
	return result
}

func GetValidatorErrMsg(e validator.ValidationErrors) string {
	var strArr []string
	for _, v := range e {
		newstr := fmt.Sprintf("%s %s", v.Field(), UnmappedrrMsg[v.Tag()])
		strArr = append(strArr, newstr)
	}

	return strings.Join(strArr, ", ")
}

// validate custome check alphanumeric only
func CustomValidationAlphanumeric(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if alphaNumeric.MatchString(value) {
		return true
	}
	return false
}

// validate custome check extention
func CustomValidationExtention(fl validator.FieldLevel) bool {

	files := fl.Field().Interface().([]*multipart.FileHeader)
	allowedExtensions := []string{".txt", ".pdf"}

	for _, file := range files {
		// File is optional, so it can be empty
		if len(strings.TrimSpace(file.Filename)) == 0 {
			continue
		}

		// // Check if the file has a valid extension
		validExtension := false
		for _, ext := range allowedExtensions {
			if strings.HasSuffix(file.Filename, ext) {
				validExtension = true
				break
			}
		}

		if !validExtension {
			return false
		}
	}

	return true
}

// validate custome check extention
func CustomValidationEnumFilter(fl validator.FieldLevel) bool {

	filter := fl.Field().String()
	allowedFilter := []string{"title", "description"}

	validFilter := false
	for _, ftr := range allowedFilter {
		if ftr == filter {
			validFilter = true
			break
		}
	}

	return validFilter
}
