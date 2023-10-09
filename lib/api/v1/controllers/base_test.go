package controllers_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BaseController", func() {
	Describe("#GetCurrentUser", func() {
		Describe("given an authenticated request", func() {
			It("returns a user", func() {
				user := fabricators.FabricateTester()
				ctx, _ := test.MakeAuthenticatedRequest(http.MethodGet, "/health-token", nil, nil, user)
				healthController := controllers.HealthController{}

				currentUser := healthController.GetCurrentUser(ctx)

				Expect(currentUser).NotTo(BeNil())
				Expect(currentUser.Base.ID).To(Equal(user.Base.ID))
			})
		})

		Describe("given a normal request", func() {
			It("does NOT return any users", func() {
				ctx, _ := test.MakeRequest(http.MethodGet, "/health-token", nil, nil)
				healthController := controllers.HealthController{}

				currentUser := healthController.GetCurrentUser(ctx)

				Expect(currentUser).To(BeNil())
			})
		})
	})
})
