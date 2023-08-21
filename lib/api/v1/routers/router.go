package routers

import (
	"github.com/markgravity/golang-ic/lib/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func ComebineRoutes(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")

	v1.GET("/health", controllers.HealthController{}.HealthStatus)
}
