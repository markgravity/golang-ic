package validators

import (
	"github.com/markgravity/golang-ic/helpers/log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	et "github.com/go-playground/validator/v10/translations/en"
)

var translator ut.Translator

func Init() {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatal("Fail to get validator")
	}

	err := registerTranslations(validate)
	if err != nil {
		log.Fatal("Fail to get register translations")
	}
}

func GetTranslator() ut.Translator {
	return translator
}

func Validate(i interface{}) error {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatal("Fail to get validator")
	}

	return validate.Struct(i)
}

func registerTranslations(validate *validator.Validate) (err error) {
	enLocale := en.New()
	uTrans := ut.New(enLocale)
	translator, _ = uTrans.GetTranslator(enLocale.Locale())

	// Register default translation
	err = et.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		return err
	}

	// Register our translations
	for _, t := range translations {
		if t.customTransFunc != nil && t.customRegisFunc != nil {
			err = validate.RegisterTranslation(t.tag, translator, t.customRegisFunc, t.customTransFunc)
		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
			err = validate.RegisterTranslation(t.tag, translator, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
			err = validate.RegisterTranslation(t.tag, translator, t.customRegisFunc, translateFunc)
		} else {
			err = validate.RegisterTranslation(t.tag, translator, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		return ut.Add(tag, translation, override)
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	translation, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return translation
}
