package helpers_test

import (
	"github.com/markgravity/golang-ic/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query Params Helpers", func() {
	Describe(".GenerateURLParams", func() {
		Context("given a map of query values", func() {
			It("returns a string of query values", func() {
				queryValues := map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				}

				result := helpers.GenerateURLParams(queryValues)

				Expect(result.Encode()).To(Equal("key1=value1&key2=value2"))
			})
		})
	})
})
