package validators

import (
	"github.com/go-playground/validator/v10"
)

var translations []struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}
