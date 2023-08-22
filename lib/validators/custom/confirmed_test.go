package custom_test

import (
	"github.com/markgravity/golang-ic/lib/validators"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#ConfirmedValidator", func() {
	Context("given VALID payload", func() {
		It("does NOT return any errors", func() {
			validators.Init()
			payload := struct {
				Password             string `binding:"confirmed"`
				PasswordConfirmation string
			}{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			}
			err := validators.Validate(payload)

			Expect(err).To(BeNil())
		})
	})

	Context("given INVALID payload", func() {
		It("returns error", func() {
			validators.Init()
			payload := struct {
				Password             string `binding:"confirmed"`
				PasswordConfirmation string
			}{
				Password:             "12345678",
				PasswordConfirmation: "INVALID",
			}
			err := validators.Validate(payload)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("given NO password confirmation", func() {
		It("returns error", func() {
			validators.Init()
			payload := struct {
				Password             string `binding:"confirmed"`
				PasswordConfirmation string
			}{
				Password: "12345678",
			}
			err := validators.Validate(payload)

			Expect(err).ToNot(BeNil())
		})
	})
})
