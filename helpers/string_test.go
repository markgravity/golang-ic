package helpers_test

import (
	"github.com/markgravity/golang-ic/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Helpers", func() {
	Describe(".ToSnakeCase", func() {
		Context("given a kebab case string", func() {
			It("returns snake case string", func() {
				kebebCaseStr := "kebab-case"
				result := helpers.ToSnakeCase(kebebCaseStr)

				Expect(result).To(Equal("kebab_case"))
			})
		})

		Context("given a camel case string", func() {
			It("returns snake case string", func() {
				camelCaseStr := "camelCase"
				result := helpers.ToSnakeCase(camelCaseStr)

				Expect(result).To(Equal("camel_case"))
			})
		})

		Context("given a pascal case string", func() {
			It("returns snake case string", func() {
				pascalCaseStr := "PascalCase"
				result := helpers.ToSnakeCase(pascalCaseStr)

				Expect(result).To(Equal("pascal_case"))
			})
		})

		Context("given a sentense case string", func() {
			It("returns snake case string", func() {
				sentenseCaseStr := "sentense case"
				result := helpers.ToSnakeCase(sentenseCaseStr)

				Expect(result).To(Equal("sentense_case"))
			})
		})

		Context("given an upper case string", func() {
			It("returns snake case string", func() {
				upperCaseStr := "UPPER CASE"
				result := helpers.ToSnakeCase(upperCaseStr)

				Expect(result).To(Equal("upper_case"))
			})
		})
	})
})
