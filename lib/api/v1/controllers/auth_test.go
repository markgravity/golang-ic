package controllers_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthController", func() {
	Describe("POST /sign-in", func() {
		Context("Given VALID payload", func() {
			It("returns status OK", func() {
				test.CleanUpDatabase()
				payload := map[string]interface{}{
					"email":                 "test@gmail.com",
					"password":              "test123",
					"password_confirmation": "test123",
				}

				ctx, resp := test.MakeRequest(http.MethodPost, "/auth/sign-in", nil, payload)
				controller := controllers.AuthController{}

				controller.SignUp(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		Context("Given INVALID payload", func() {
			It("returns an unprocessable entity status", func() {
				test.CleanUpDatabase()
				payload := map[string]interface{}{
					"email":                 "INVALID",
					"password":              "test123",
					"password_confirmation": "test123",
				}

				ctx, resp := test.MakeRequest(http.MethodPost, "/auth/sign-in", nil, payload)
				controller := controllers.AuthController{}

				controller.SignUp(ctx)

				Expect(resp.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})

		Context("Given duplicated email", func() {
			It("returns an unprocessable entity status", func() {
				test.CleanUpDatabase()
				form := forms.SignUpForm{
					Email:                "test@gmail.com",
					Password:             "test123",
					PasswordConfirmation: "test123",
				}
				_, _ = form.Save()

				payload := map[string]interface{}{
					"email":                 "test@gmail.com",
					"password":              "test123",
					"password_confirmation": "test123",
				}

				ctx, resp := test.MakeRequest(http.MethodPost, "/auth/sign-in", nil, payload)
				controller := controllers.AuthController{}

				controller.SignUp(ctx)

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
