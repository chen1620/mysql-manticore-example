package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"mysql-manticore-example/configs"
	"mysql-manticore-example/controllers"
	"mysql-manticore-example/databases"
	"mysql-manticore-example/routers"
)

func main() {
	r := gin.Default()

	appCfg := configs.LoadAppConfig()

	mysqlDB := databases.InitializeMySQL(appCfg.MySQLConfig)
	manticoreClient := databases.InitializeManticore(appCfg.ManticoreConfig)

	var (
		userController controllers.UserController
		postController controllers.PostController
	)

	{
		repo := databases.NewUserRepository(mysqlDB)
		userController = controllers.NewUserController(repo)
	}

	{
		repo := databases.NewPostRepository(mysqlDB, manticoreClient)
		postController = controllers.NewPostController(repo)
	}

	routes := routers.NewRoutes(
		routers.NewUserRoutes(r, userController),
		routers.NewPostRoutes(r, postController),
	)

	routes.Setup()

	s := &http.Server{
		Addr:           appCfg.HTTPAddress,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %+v", err)
	}
}
