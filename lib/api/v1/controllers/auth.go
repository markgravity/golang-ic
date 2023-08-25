package controllers

import (
	"net/http"

	jsonhelpers "github.com/markgravity/golang-ic/helpers/json"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"

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
