package forms_test

import (
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/validators"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keywords", func() {
	Describe("Save", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given VALID form", func() {
			It("returns without error", func() {
				_, fileHeader, _ := test.GetMultipartAttributesFromFile(
					"keywords/valid.csv",
					"text/csv",
				)
				user := fabricators.FabricateUser("test@gmail.com", "12345678")
				form := forms.KeywordsForm{
					FileHeader: fileHeader,
					User:       user,
				}

				err := form.Save()

				Expect(err).To(BeNil())
			})
		})

		Context("Given INVALID file type", func() {
			It("returns an error", func() {
				user := fabricators.FabricateUser("test@gmail.com", "12345678")
				_, fileHeader, _ := test.GetMultipartAttributesFromFile(
					"keywords/text.txt",
					"text/plain",
				)
				form := forms.KeywordsForm{
					FileHeader: fileHeader,
					User:       user,
				}

				err := form.Save()

				Expect(err.Error()).To(Equal("file type is not supported"))
			})
		})

		Context("Given INVALID file content", func() {
			It("returns an error", func() {
				user := fabricators.FabricateUser("test@gmail.com", "12345678")
				_, fileHeader, _ := test.GetMultipartAttributesFromFile(
					"keywords/empty.csv",
					"text/csv",
				)
				form := forms.KeywordsForm{
					FileHeader: fileHeader,
					User:       user,
				}

				err := form.Save()

				Expect(err.Error()).To(Equal("CSV file only accepts from 1 to 1000 keywords"))
			})
		})

		Context("Given NULL file", func() {
			It("returns an error", func() {
				user := fabricators.FabricateUser("test@gmail.com", "12345678")
				form := forms.KeywordsForm{
					FileHeader: nil,
					User:       user,
				}

				err := validators.Validate(form)
				validationErrors := err.(validator.ValidationErrors)
				Expect(validationErrors).To(HaveLen(1))
				Expect(validationErrors[0].Tag()).To(Equal("required"))
			})
		})
	})
})
