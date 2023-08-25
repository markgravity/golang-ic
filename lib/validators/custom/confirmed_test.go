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
				Field             string `binding:"confirmed"`
				FieldConfirmation string
			}{
				Field:             "12345678",
				FieldConfirmation: "12345678",
			}
			err := validators.Validate(payload)

			Expect(err).To(BeNil())
		})
	})

	Context("given INVALID payload", func() {
		It("returns error", func() {
			validators.Init()
			payload := struct {
				Field             string `binding:"confirmed"`
				FieldConfirmation string
			}{
				Field:             "12345678",
				FieldConfirmation: "INVALID",
			}
			err := validators.Validate(payload)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("given NO field confirmation", func() {
		It("returns error", func() {
			validators.Init()
			payload := struct {
				Field             string `binding:"confirmed"`
				FieldConfirmation string
			}{
				Field: "12345678",
			}
			err := validators.Validate(payload)

			Expect(err).ToNot(BeNil())
		})
	})
})
