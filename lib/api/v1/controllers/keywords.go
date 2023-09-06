package controllers

import (
	"net/http"

	jsonhelpers "github.com/markgravity/golang-ic/helpers/json"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/gin-gonic/gin"
)

type KeywordsController struct {
	BaseController
}

func (KeywordsController) Upload(ctx *gin.Context) {
	form := forms.KeywordsForm{}

	err := ctx.ShouldBind(&form)
	if err != nil {
		jsonhelpers.RenderErrorWithDefaultCode(ctx, http.StatusBadRequest, err)
		return
	}

	value, _ := ctx.Get(UserKey)
	user := value.(models.User)
	form.User = &user

	err = form.Save()
	if err != nil {
		jsonhelpers.RenderUnprocessableEntityError(ctx, err)
		return
	}

	jsonhelpers.RenderJSON(ctx, http.StatusOK, nil)
}