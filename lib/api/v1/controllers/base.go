package controllers

import (
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (c *BaseController) GetCurrentUser(ctx *gin.Context) *models.User {
	value, _ := ctx.Get(UserKey)
	if value == nil {
		return nil
	}

	user := value.(models.User)

	return &user
}
