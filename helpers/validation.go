package helpers

import (
	"github.com/beego/beego/v2/core/validation"
)

func GetValidationResult(valid validation.Validation) map[string]string {
	validationResult := make(map[string]string)

	for _, err := range valid.Errors {
		validationResult[err.Key] = err.Message
	}
	return validationResult
}
