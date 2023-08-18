package custom

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ConfirmedValidator(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field().String()
	fieldConfirmationName := fmt.Sprintf("%vConfirmation", fieldLevel.FieldName())
	fieldConfirmation := fieldLevel.Top().FieldByName(fieldConfirmationName).String()

	return field == fieldConfirmation
}
