package server

import (
	"crudgin/api/handlers"
	"crudgin/api/handlers/middleware"
	"crudgin/entities/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Hi"})
	})
	userRoute := route.Group("/user")
	userRoute.Use(middleware.AuthMiddleware())

	route.GET("/products", handlers.GetAllProduct)
	route.POST("/products", handlers.CreateProduct)
	route.GET("/products/:id", handlers.GetProduct)
	route.PATCH("/products/:id", handlers.UpdateProduct)
	route.DELETE("/products/:id", handlers.DeleteProduct)

	route.POST("/createuser", handlers.CreateUser)
	route.POST("/login", user.Login)
	userRoute.GET("/users/:id", handlers.GetAllUser)
	userRoute.GET("/user/:id", handlers.GetUser)
	userRoute.PUT("/user/:id", handlers.UpdateUser)
	userRoute.DELETE("/user/:id", handlers.DeleteUSer)

	route.Run()

}
