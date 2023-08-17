package oauth_test

import (
	"net/http"

	"github.com/markgravity/golang-ic/lib/services/oauth"
	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthServer", func() {
	Describe("HandleTokenRequest", func() {
		Context("Given VALID params", func() {
			It("returns status OK", func() {
				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      "test",
					"password":      "test",
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			It("returns correct response body", func() {
				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "1",
					"client_secret": "2",
					"username":      "test",
					"password":      "test",
				}
				ctx, resp := test.MakePostFormRequest("/oauth", headers, params)

				_ = oauth.SetUpOAuthServer()
				server := oauth.GetOAuthServer()
				_ = server.HandleTokenRequest(ctx.Writer, ctx.Request)

				Expect(resp.Result()).To(test.MatchJSONSchema("token/success"))
			})
		})

		Context("Given INVALID params", func() {
			It("returns error", func() {
				headers := map[string]string{
					"Content-Type": "multipart/form-data",
				}
				params := map[string]interface{}{
					"grant_type":    "password",
					"client_id":     "INVALID",
					"client_secret": "2",
					"username":      "test",
					"password":      "test",
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
