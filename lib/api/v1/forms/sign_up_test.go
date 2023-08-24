package forms_test

import (
	"github.com/markgravity/golang-ic/helpers"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"

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
})
