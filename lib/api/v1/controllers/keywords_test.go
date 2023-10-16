package controllers_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/lib/middlewares"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	"github.com/gin-gonic/gin"
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

	Describe("GET /keywords", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})
		Context("Given VALID payload", func() {
			It("returns the status OK", func() {
				user := fabricators.FabricateTester()
				payload := map[string]interface{}{
					"limit": "1",
				}
				ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords", nil, payload, user)

				controller := controllers.KeywordsController{}
				controller.Index(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			It("returns the keywords that belong to the current user, serialized in JSON", func() {
				user := fabricators.FabricateTester()
				fabricators.FabricateKeyword("k1", user)
				fabricators.FabricateKeyword("k2", user)

				payload := map[string]interface{}{
					"limit": "1",
				}
				ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords", nil, payload, user)

				controller := controllers.KeywordsController{}
				controller.Index(ctx)

				Expect(resp.Result()).To(test.MatchJSONSchema("response/keyword/success"))
			})
		})

		Context("Given INVALID payload", func() {
			Context("Given NO limit", func() {
				It("returns the unprocessable entity status", func() {
					user := fabricators.FabricateTester()
					ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords", nil, nil, user)

					controller := controllers.KeywordsController{}
					controller.Index(ctx)

					Expect(resp.Code).To(Equal(http.StatusUnprocessableEntity))
				})
			})

			Context("Given INVALID offset", func() {
				It("returns the bad request status", func() {
					user := fabricators.FabricateTester()
					payload := map[string]interface{}{
						"limit":  "1",
						"offset": "INVALID",
					}
					ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords", nil, payload, user)

					controller := controllers.KeywordsController{}
					controller.Index(ctx)

					Expect(resp.Code).To(Equal(http.StatusBadRequest))
				})
			})
		})
	})

	Describe("GET /keyword/:id", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})
		Context("Given VALID id", func() {
			It("returns the status OK", func() {
				user := fabricators.FabricateTester()
				keyword := fabricators.FabricateKeyword("k1", user)
				ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords/"+keyword.Base.ID.String(), nil, nil, user)
				ctx.Params = []gin.Param{
					{
						Key:   "id",
						Value: keyword.Base.ID.String(),
					},
				}
				controller := controllers.KeywordsController{}
				controller.Show(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			It("returns the keyword that belong to the current user, serialized in JSON", func() {
				user := fabricators.FabricateTester()
				keyword := fabricators.FabricateKeyword("k1", user)
				ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords/"+keyword.Base.ID.String(), nil, nil, user)
				ctx.Params = []gin.Param{
					{
						Key:   "id",
						Value: keyword.Base.ID.String(),
					},
				}

				controller := controllers.KeywordsController{}
				controller.Show(ctx)

				Expect(resp.Result()).To(test.MatchJSONSchema("response/keyword_detail/success"))
			})
		})

		Context("Given INVALID id", func() {
			It("returns the unprocessable entity status", func() {
				user := fabricators.FabricateTester()
				ctx, resp := test.MakeAuthenticatedRequest(http.MethodGet, "/keywords/invalid", nil, nil, user)
				ctx.Params = []gin.Param{
					{
						Key:   "id",
						Value: "INVALID",
					},
				}

				controller := controllers.KeywordsController{}
				controller.Show(ctx)

				Expect(resp.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})
	})
})
