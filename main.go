package main

import (
	"fmt"
	"thecrafterja/gingormtut/controllers"
	"thecrafterja/gingormtut/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//Tutorial: https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
	fmt.Println("Running...")

	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}
