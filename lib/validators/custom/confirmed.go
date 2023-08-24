package custom

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ConfirmedValidator(fieldLevel validator.FieldLevel) bool {
	fieldValue := fieldLevel.Field().String()
	fieldConfirmationName := fmt.Sprintf("%vConfirmation", fieldLevel.FieldName())
	fieldConfirmationValue := fieldLevel.Top().FieldByName(fieldConfirmationName).String()

	return fieldValue == fieldConfirmationValue
}
