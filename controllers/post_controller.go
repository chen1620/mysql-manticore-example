package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mysql-manticore-example/databases"
	"mysql-manticore-example/models"
)

// PostController data type.
type PostController struct {
	repo databases.PostRepository
}

// NewPostController creates new post controller.
func NewPostController(repo databases.PostRepository) PostController {
	return PostController{
		repo: repo,
	}
}

// GetPostByID gets post by id.
func (p *PostController) GetPostByID(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	postDetail, err := p.repo.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": postDetail,
	})
}

func (p *PostController) AddPost(c *gin.Context) {
	// todo: shouldn't use model here, define other struct to bind request.
	post := models.Post{}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := p.repo.AddPost(c, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": post})
}
