package validators_test

import (
	"github.com/markgravity/golang-ic/lib/validators"

	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("validators", func() {
	Describe("#Validate", func() {
		Context("given a built-in rule", func() {
			It("returns a correct error message", func() {
				validators.Init()

				payload := struct {
					Module string `binding:"required"`
				}{}

				err := validators.Validate(payload).(validator.ValidationErrors)

				Expect(err[0].Translate(validators.GetTranslator())).To(Equal("Module is a required field"))
			})
		})

		Context("given a custom rule", func() {
			It("returns a correct error message", func() {
				validators.Init()

				payload := struct {
					Module string `binding:"validModule"`
				}{
					Module: "INVALID",
				}

				err := validators.Validate(payload).(validator.ValidationErrors)

				Expect(err[0].Translate(validators.GetTranslator())).To(Equal("Module is invalid"))
			})
		})
	})

	Describe("#GetTranslator", func() {
		It("returns a translator", func() {
			validators.Init()
			translator := validators.GetTranslator()

			Expect(translator).ToNot(BeNil())
		})
	})
})
