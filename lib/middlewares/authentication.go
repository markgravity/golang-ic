package middlewares

import (
	"net/http"
	"strings"

	"github.com/markgravity/golang-ic/database"
	jsonhelpers "github.com/markgravity/golang-ic/helpers/json"
	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

func HandleAuthenticatedRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. Get access token from header
		accessTokenForm := forms.AccessTokenForm{}
		err := ctx.ShouldBindHeader(&accessTokenForm)
		if err != nil {
			jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusUnauthorized, err)
			return
		}

		// 2. Get user id from access token
		token := getTokenFromAccessToken(accessTokenForm.AccessToken)
		authServer := oauth.GetOAuthServer()
		tokenInfo, err := authServer.Manager.LoadAccessToken(ctx, token)
		if err != nil {
			jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusUnauthorized, err)
			return
		}

		// 3. Set user to context
		db := database.GetDB()
		var user models.User
		err = db.First(&user, "id = ?", tokenInfo.GetUserID()).Error
		if err != nil {
			jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusUnauthorized, err)
			return
		}

		ctx.Set(controllers.UserKey, user)

		ctx.Next()
	}
}

func getTokenFromAccessToken(accessToken string) string {
	parts := strings.Split(accessToken, " ")
	if len(parts) <= 1 {
		return ""
	}

	return parts[1]
}
