package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/markgravity/golang-ic/lib/api/v1/controllers"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.HealthController{}.HealthStatus)

	v1.POST("/auth/sign-up", controllers.AuthController{}.SignUp)
}
