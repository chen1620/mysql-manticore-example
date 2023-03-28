package controllers

import (
	"github.com/gin-gonic/gin"
	"mysql-manticore-example/databases"
	"net/http"
)

// UserController data type.
type UserController struct {
	repo databases.UserRepository
}

// NewUserController ...
func NewUserController(repo databases.UserRepository) UserController {
	return UserController{
		repo: repo,
	}
}

// Login ...
func (u *UserController) Login(c *gin.Context) {
	// implement logic here ...
	c.JSON(http.StatusOK, gin.H{})
}
