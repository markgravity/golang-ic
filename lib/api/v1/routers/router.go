package routers

import (
	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
	"github.com/markgravity/golang-ic/lib/middlewares"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.HealthController{}.HealthStatus)

	v1.POST("/auth/sign-in", controllers.AuthController{}.SignIn)

	v1.POST("/auth/sign-up", controllers.AuthController{}.SignUp)

	// Authenticated routes
	authenticatedV1 := v1
	authenticatedV1.Use(middlewares.HandleAuthenticatedRequest())

	authenticatedV1.POST("/keywords/upload", controllers.KeywordsController{}.Upload)
}
