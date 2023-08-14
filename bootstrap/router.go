package bootstrap

import (
	apiv1router "github.com/nimblehq/mark-ic/lib/api/v1/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiv1router.ComebineRoutes(r)

	return r
}