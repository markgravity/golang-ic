package jsonhelpers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	jsonhelpers "github.com/markgravity/golang-ic/helpers/json"
	. "github.com/markgravity/golang-ic/test"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSON Response Helpers", func() {
	Context(".RenderErrorWithDefaultCode", func() {
		Context("given status code is 400", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusBadRequest, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusBadRequest)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("bad_request_error"))
			})
		})

		Context("given status code is 403", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusForbidden, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusForbidden)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("forbidden_error"))
			})
		})

		Context("given status code is 500", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusInternalServerError, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusInternalServerError)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("internal_server_error"))
			})
		})

		Context("given status code is 422", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusUnprocessableEntity, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusUnprocessableEntity)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("unprocessable_entity_error"))
			})
		})

		Context("given status code is 404", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusNotFound, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusNotFound)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("not_found_error"))
			})
		})

		Context("given status code is 401", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusUnauthorized, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusUnauthorized)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("unauthorized_error"))
			})
		})

		Context("given status code is 422", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusUnprocessableEntity, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusUnprocessableEntity)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("unprocessable_entity_error"))
			})
		})

		Context("given status code is 503", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusServiceUnavailable, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusServiceUnavailable)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal("service_unavailable_error"))
			})
		})

		Context("given status code is not defined", func() {
			It("renders the error payload with the default error code", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()

				jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusBadGateway, errors.New(errorMessage))

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusBadGateway)))
				Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
				Expect(errorResponse.Errors[0].Code).To(Equal(http.StatusText(http.StatusBadGateway)))
			})
		})
	})

	Context(".RenderError", func() {
		It("renders error with the given status", func() {
			responseRecorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseRecorder)

			jsonhelpers.RenderError(ctx, http.StatusBadRequest, faker.Sentence(), faker.Word())

			Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
		})

		It("renders the error payload with the given details", func() {
			responseRecorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseRecorder)
			errorMessage := faker.Sentence()
			errorCode := faker.Word()

			jsonhelpers.RenderError(ctx, http.StatusBadRequest, errorMessage, errorCode)

			Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

			errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
			Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusBadRequest)))
			Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
			Expect(errorResponse.Errors[0].Code).To(Equal(errorCode))
		})
	})

	Context(".RenderErrorWithMeta", func() {
		It("renders error with the given status", func() {
			responseRecorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseRecorder)
			meta := &map[string]interface{}{"Message": "message"}

			jsonhelpers.RenderErrorWithMeta(ctx, http.StatusBadRequest, faker.Sentence(), faker.Word(), meta)

			Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
		})

		It("renders the error payload with the given details", func() {
			responseRecorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(responseRecorder)
			errorMessage := faker.Sentence()
			errorCode := faker.Word()
			meta := &map[string]interface{}{"message": "message"}

			jsonhelpers.RenderErrorWithMeta(ctx, http.StatusBadRequest, errorMessage, errorCode, meta)

			Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

			errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
			Expect(errorResponse.Errors[0].Title).To(Equal(http.StatusText(http.StatusBadRequest)))
			Expect(errorResponse.Errors[0].Detail).To(Equal(errorMessage))
			Expect(errorResponse.Errors[0].Code).To(Equal(errorCode))
			Expect(errorResponse.Errors[0].Meta).To(Equal(meta))
		})

		Context("given meta data keys are NOT snake case", func() {
			It("renders the error meta data keys in snake case", func() {
				responseRecorder := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(responseRecorder)
				errorMessage := faker.Sentence()
				errorCode := faker.Word()
				meta := &map[string]interface{}{
					"kebab-case-key": "value",
					"PascalCaseKey":  "value",
					"camelCaseKey":   "value",
					"snake_case_key": "value",
				}
				expectedMeta := &map[string]interface{}{
					"kebab_case_key":  "value",
					"pascal_case_key": "value",
					"camel_case_key":  "value",
					"snake_case_key":  "value",
				}

				jsonhelpers.RenderErrorWithMeta(ctx, http.StatusBadRequest, errorMessage, errorCode, meta)

				Expect(responseRecorder.Result()).To(MatchJSONSchema("error"))

				errorResponse := UnmarshalErrorResponseBody(responseRecorder.Body)
				Expect(errorResponse.Errors[0].Meta).To(Equal(expectedMeta))
			})
		})
	})
})
