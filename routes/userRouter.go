package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ramdhanyA/task-5-vix-btpns"
	"github.com/ramdhanyA/task-5-vix-btpns/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
