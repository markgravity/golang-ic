package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/services/oauth"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.HealthController{}.HealthStatus)

	// TODO: Correct the logic in Sign-In task
	v1.POST("/auth/sign-in", func(context *gin.Context) {
		server := oauth.GetOAuthServer()
		err := server.HandleTokenRequest(context.Writer, context.Request)

		if err != nil {
			_ = context.AbortWithError(403, err)
		}
	})
}
