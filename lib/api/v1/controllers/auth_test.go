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
	Describe("POST /sign-up", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given VALID payload", func() {
			It("returns the status OK", func() {
				payload := map[string]interface{}{
					"email":                 "test@gmail.com",
					"password":              "test123",
					"password_confirmation": "test123",
				}

				ctx, resp := test.MakeRequest(http.MethodPost, "/auth/sign-up", nil, payload)
				controller := controllers.AuthController{}

				controller.SignUp(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		Context("Given INVALID payload", func() {
			It("returns the unprocessable entity status", func() {
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

				ctx, resp := test.MakeRequest(http.MethodPost, "/auth/sign-up", nil, payload)
				controller := controllers.AuthController{}

				controller.SignUp(ctx)

				Expect(resp.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})
	})

	Describe("POST /sign-in", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given VALID payload", func() {
			It("returns status OK", func() {
				form := forms.SignUpForm{
					Email:                "test@gmail.com",
					Password:             "test123",
					PasswordConfirmation: "test123",
				}
				_, _ = form.Save()

				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				payload := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      form.Email,
					"password":      form.Password,
				}

				ctx, resp := test.MakePostFormRequest("/auth/sign-in", headers, payload)
				controller := controllers.AuthController{}

				controller.SignIn(ctx)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			It("returns correct response body", func() {
				form := forms.SignUpForm{
					Email:                "test@gmail.com",
					Password:             "test123",
					PasswordConfirmation: "test123",
				}
				_, _ = form.Save()

				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				payload := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      form.Email,
					"password":      form.Password,
				}

				ctx, resp := test.MakePostFormRequest("/auth/sign-in", headers, payload)
				controller := controllers.AuthController{}

				controller.SignIn(ctx)

				Expect(resp.Result()).To(test.MatchJSONSchema("response/token/success"))
			})
		})

		Context("Given INVALID payload", func() {
			It("returns the bad request status", func() {
				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				payload := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      "INVALID",
					"password":      "123",
				}

				ctx, resp := test.MakePostFormRequest("/auth/sign-in", headers, payload)
				controller := controllers.AuthController{}

				controller.SignIn(ctx)

				Expect(resp.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
