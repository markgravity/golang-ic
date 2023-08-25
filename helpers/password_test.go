package helpers_test

import (
	"github.com/markgravity/golang-ic/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Password", func() {
	Describe("HashPassword", func() {
		Context("given a string", func() {
			It("returns the hashed string", func() {
				hashedPassword, _ := helpers.HashPassword("hello password")

				Expect(len(hashedPassword) > 1).To(BeTrue())
				Expect(hashedPassword).To(ContainSubstring("$"))
			})
		})
	})

	Describe("ComparePassword", func() {
		Context("given a correct password", func() {
			It("returns nil", func() {
				hashedPassword := "$2a$10$9rnyG9shT0T49i9CBQcnYuICoTwgKvwFCY/EEET63PqjJTat1qHRW"
				correctPassword := "123456"
				Expect(helpers.ComparePassword(hashedPassword, correctPassword)).To(BeNil())
			})
		})

		Context("given an incorrect password", func() {
			It("is NOT nil", func() {
				hashedPassword := "$2a$10$9rnyG9shT0T49i9CBQcnYuICoTwgKvwFCY/EEET63PqjJTat1qHRW"
				incorrectPassword := "111111"
				Expect(helpers.ComparePassword(hashedPassword, incorrectPassword)).NotTo(BeNil())
			})
		})
	})
})
