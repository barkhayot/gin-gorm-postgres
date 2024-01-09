package main

import (
	"fmt"
	"go-curd/controllers"
	"go-curd/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
}

func main() {
	fmt.Println("hello")

	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PUT("/posts/:id", controllers.UpdatePostById)
	r.DELETE("/posts/:id", controllers.DeletePostById)
	r.Run()
}
