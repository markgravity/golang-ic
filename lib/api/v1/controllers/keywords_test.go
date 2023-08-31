package controllers_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordsController", func() {
	Describe("POST /keywords/upload", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given VALID payload", func() {
			It("returns the status OK", func() {
				fabricators.FabricateUser("test@gmail.com", "123456")
				ctx, resp := test.MakeMultipartRequestRequest("/keywords/upload", "keywords/valid.csv", "text/csv")

				controller := controllers.KeywordsController{}
				controller.Upload(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
