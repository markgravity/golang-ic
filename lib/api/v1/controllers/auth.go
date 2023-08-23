package controllers

import (
	"net/http"

	jsonhelpers "github.com/markgravity/golang-ic/helpers/json"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/api/v1/serializers"
	"github.com/markgravity/golang-ic/lib/services/oauth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

func (AuthController) SignUp(ctx *gin.Context) {
	form := forms.SignUpForm{}

	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = form.Save()
	if err != nil {
		jsonhelpers.RenderUnprocessableEntityError(ctx, err)
		return
	}

	jsonhelpers.RenderJSON(ctx, http.StatusOK, nil)
}

func (AuthController) SignIn(ctx *gin.Context) {
	server := oauth.GetOAuthServer()

	gt, tgr, err := server.ValidationTokenRequest(ctx.Request)
	if err != nil {
		jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusBadRequest, err)
		return
	}

	token, err := server.GetAccessToken(ctx, gt, tgr)
	if err != nil {
		jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusBadRequest, err)
		return
	}

	serializer := serializers.TokenSerializer{
		Token:     token,
		TokenType: server.Config.TokenType,
	}

	jsonhelpers.RenderJSON(ctx, http.StatusOK, serializer.Data())
}
