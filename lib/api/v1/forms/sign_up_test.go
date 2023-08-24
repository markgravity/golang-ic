package forms_test

import (
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/validators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SignUp", func() {
	Describe("Save", func() {
		Context("Given VALID form", func() {
			It("returns without error", func() {
				form := forms.SignUpForm{
					Email:    "example@gmail.com",
					Password: "12345678",
				}

				_, err := form.Save()

				Expect(err).To(BeNil())
			})

			It("returns correct user", func() {
				form := forms.SignUpForm{
					Email:    "example2@gmail.com",
					Password: "12345678",
				}

				user, _ := form.Save()

				Expect(user.Base.ID).ToNot(BeNil())
				Expect(user.Email).To(Equal(form.Email))
				Expect(helpers.ComparePassword(user.EncryptedPassword, form.Password)).To(BeNil())
			})
		})

		Context("Given duplicated email", func() {
			It("returns an error", func() {
				form := forms.SignUpForm{
					Email:    "example2@gmail.com",
					Password: "12345678",
				}

				_, _ = form.Save()
				_, err := form.Save()

				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("Validations", func() {
		Context("Given INVALID email", func() {
			It("returns the error", func() {
				form := forms.SignUpForm{
					Email:                "INVALID",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				}

				err := validators.Validate(form)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("Given EMPTY password", func() {
			It("returns the error", func() {
				form := forms.SignUpForm{
					Email:                "example@gmail.com",
					PasswordConfirmation: "12345678",
				}

				err := validators.Validate(form)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("Given password less than min requirement", func() {
			It("returns the error", func() {
				form := forms.SignUpForm{
					Email:                "example@gmail.com",
					Password:             "123",
					PasswordConfirmation: "123",
				}

				err := validators.Validate(form)

				Expect(err).ToNot(BeNil())
			})
		})

		Context("Given INVALID password confirmation", func() {
			It("returns the error", func() {
				form := forms.SignUpForm{
					Email:                "example@gmail.com",
					Password:             "12345678",
					PasswordConfirmation: "INVALID",
				}

				err := validators.Validate(form)

				Expect(err).ToNot(BeNil())
			})
		})
	})
})
