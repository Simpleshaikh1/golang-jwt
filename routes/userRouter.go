package routes

import (
	controller "github.com/Simpleshaikh1/golang-jwt/controllers"
	"github.com/Simpleshaikh1/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
}
