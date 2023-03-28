package routers

import (
	"github.com/gin-gonic/gin"

	"mysql-manticore-example/controllers"
)

// UserRoutes struct.
type UserRoutes struct {
	handler        *gin.Engine
	userController controllers.UserController
}

// Setup user routes
func (u UserRoutes) Setup() {
	api := u.handler.Group("/api/users")
	{
		api.POST("/login", u.userController.Login)
	}
}

// NewUserRoutes creates user routes.
func NewUserRoutes(handler *gin.Engine, controller controllers.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: controller,
	}
}
