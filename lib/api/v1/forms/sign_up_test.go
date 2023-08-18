package forms_test

import (
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SignUp", func() {
	Describe("Valid", func() {
		Context("Given VALID form", func() {
			It("returns without error", func() {
				form := forms.SignUpForm{
					Email:    "example@gmail.com",
					Password: "12345678",
				}

				err := form.Validate()

				Expect(err).To(BeNil())
			})
		})

		Context("Given INVALID email", func() {
			It("returns error", func() {
				form := forms.SignUpForm{
					Email:    "INVALID",
					Password: "12345678",
				}

				err := form.Validate()

				Expect(err).ToNot(BeNil())
			})
		})

		Context("Given INVALID password", func() {
			It("returns error", func() {
				form := forms.SignUpForm{
					Email:    "example@gmail.com",
					Password: "1",
				}

				err := form.Validate()

				Expect(err).ToNot(BeNil())
			})
		})
	})

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
					Email:    "example@gmail.com",
					Password: "12345678",
				}

				user, _ := form.Save()

				Expect(user.Base.ID).ToNot(BeNil())
				Expect(user.Email).To(Equal(form.Email))
				Expect(helpers.ComparePassword(user.EncryptedPassword, form.Password)).To(BeNil())
			})
		})
	})
})
