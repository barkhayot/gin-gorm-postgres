package main

import (
	"go-curd/initializer"
	"go-curd/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Post{})

}
