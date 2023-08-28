package oauth_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/services/oauth"
	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthServer", func() {
	Describe("HandleTokenRequest", func() {
		Context("Given VALID params", func() {
			It("responds with the OK status", func() {
				test.CleanUpDatabase()

				form := forms.SignUpForm{
					Email:                "test@gmail.com",
					Password:             "test123",
					PasswordConfirmation: "test123",
				}
				_, _ = form.Save()

				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      form.Email,
					"password":      form.Password,
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			It("returns correct response body", func() {
				test.CleanUpDatabase()

				form := forms.SignUpForm{
					Email:                "test@gmail.com",
					Password:             "test123",
					PasswordConfirmation: "test123",
				}
				_, _ = form.Save()

				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      form.Email,
					"password":      form.Password,
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Result()).To(test.MatchJSONSchema("token/success"))
			})
		})

		Context("Given INVALID client credential", func() {
			It("responds with the internal server error status code", func() {
				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "INVALID",
					"client_secret": "2",
					"username":      "test@gmail.com",
					"password":      "test123",
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("Given INVALID email", func() {
			It("responds with the internal server error status code", func() {
				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      "INVALID",
					"password":      "test123",
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("Given INVALID password", func() {
			It("responds with the internal server error status code", func() {
				test.CleanUpDatabase()

				form := forms.SignUpForm{
					Email:                "test@gmail.com",
					Password:             "test123",
					PasswordConfirmation: "test123",
				}
				_, _ = form.Save()

				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      "test@gmail.com",
					"password":      "INVALID",
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
