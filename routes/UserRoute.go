package routes;

import (
	"demo/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup) {
	userRoute := rg.Group("/users")

	userRoute.GET("", controllers.VerifyJWT, controllers.GetUsers)
}
