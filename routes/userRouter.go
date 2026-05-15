package routes

import (
	controller "github.com/akhil/golang-jwt-project/controllers"
	"github.com/akhil/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	protected := incomingRoutes.Group("/")
	protected.Use(middleware.Authenticate())
	protected.GET("/users", controller.GetUsers())
	protected.GET("/users/me", controller.GetProfile())
	protected.GET("/users/:user_id", controller.GetUser())
}
