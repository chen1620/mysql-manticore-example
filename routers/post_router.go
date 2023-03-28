package routers

import (
	"github.com/gin-gonic/gin"

	"mysql-manticore-example/controllers"
)

// PostRoutes type.
type PostRoutes struct {
	handler        *gin.Engine
	postController controllers.PostController
}

// Setup post routes.
func (p PostRoutes) Setup() {
	api := p.handler.Group("/api/posts")
	{
		api.POST("/", p.postController.AddPost)
		api.GET("/:id", p.postController.GetPostByID)
	}
}

// NewPostRoutes creates new post routes.
func NewPostRoutes(handler *gin.Engine, controller controllers.PostController) PostRoutes {
	return PostRoutes{
		handler:        handler,
		postController: controller,
	}
}
