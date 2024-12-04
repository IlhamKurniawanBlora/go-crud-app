package routes

import (
	"go-crud-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Melayani file statis
	r.Static("/public", "./public")
	r.Static("/uploads", "./uploads")

	// Route untuk melayani file index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	// Tambahkan route API CRUD
	api := r.Group("/api")
	{
		api.GET("/items", controllers.GetItems)
		api.POST("/items", controllers.CreateItem)
		api.DELETE("/items/:id", controllers.DeleteItem)
	}

	return r
}
