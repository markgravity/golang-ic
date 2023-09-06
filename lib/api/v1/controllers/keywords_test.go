package controllers_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/lib/middlewares"
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
				user := fabricators.FabricateTester()
				ctx, resp := test.MakeMultipartRequestRequest(
					"/keywords/upload",
					"keywords/valid.csv",
					"text/csv",
					nil,
					user,
				)
				middlewares.HandleAuthenticatedRequest()(ctx)

				controller := controllers.KeywordsController{}
				controller.Upload(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		Context("Given EMPTY payload", func() {
			It("returns the unprocessable entity status", func() {
				ctx, resp := test.MakeRequest(http.MethodPost, "/keywords/upload", nil, nil)

				controller := controllers.KeywordsController{}
				controller.Upload(ctx)

				Expect(resp.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})

		Context("Given INVALID payload", func() {
			It("returns the unprocessable entity status", func() {
				user := fabricators.FabricateTester()
				ctx, resp := test.MakeMultipartRequestRequest(
					"/keywords/upload",
					"keywords/invalid.csv",
					"text/csv",
					nil,
					user,
				)
				middlewares.HandleAuthenticatedRequest()(ctx)

				controller := controllers.KeywordsController{}
				controller.Upload(ctx)

				Expect(resp.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})
	})
})