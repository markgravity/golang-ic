package middlewares_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/lib/middlewares"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authentication", func() {
	AfterEach(func() {
		test.CleanUpDatabase()
	})

	Describe("HandleAuthenticatedRequest", func() {
		Context("given VALID headers", func() {
			It("returns status OK", func() {
				user := fabricators.FabricateTester()
				ctx, response := test.MakeAuthenticatedRequest(http.MethodGet, "/", nil, nil, user)

				middlewares.HandleAuthenticatedRequest()(ctx)

				Expect(response.Result().StatusCode).To(Equal(http.StatusOK))
			})

			It("sets the 'user' key with the current user id, into the context", func() {
				user := fabricators.FabricateTester()
				ctx, _ := test.MakeAuthenticatedRequest(http.MethodGet, "/", nil, nil, user)

				middlewares.HandleAuthenticatedRequest()(ctx)

				_, exists := ctx.Get(controllers.UserKey)
				Expect(exists).To(BeTrue())
			})
		})

		Context("given INVALID headers", func() {
			Context("given NO Authorization header", func() {
				It("returns the unprocessable status", func() {
					headers := map[string]string{}
					ctx, response := test.MakeRequest(http.MethodGet, "/", headers, nil)

					middlewares.HandleAuthenticatedRequest()(ctx)

					Expect(response.Result().StatusCode).To(Equal(http.StatusUnprocessableEntity))
				})
			})

			Context("given INVALID access token", func() {
				It("returns the unauthorized status", func() {
					headers := map[string]string{
						"Authorization": "Bearer INVALID",
					}

					ctx, response := test.MakeRequest(http.MethodGet, "/", headers, nil)

					middlewares.HandleAuthenticatedRequest()(ctx)

					Expect(response.Result().StatusCode).To(Equal(http.StatusUnauthorized))
				})
			})
		})
	})
})
