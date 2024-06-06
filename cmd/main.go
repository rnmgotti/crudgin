package main

import (
	"crudgin/api/server"
	"crudgin/entities/user"
	"crudgin/pkg/db"
)

func main() {
	db.ConnectDB()
	user.MigrateModels()
	server.InitRouter()
}

//route := gin.Default()
/*route.GET("/", func(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hi"})
})
userRoute := route.Group("/user")
userRoute.Use(middleware.AuthMiddleware())

route.GET("/products", product.GetAllProducts)
route.POST("/products", product.CreateProducts)
route.GET("/products/:id", product.GetProducts)
route.PATCH("/products/:id", product.UpdateProduct)
route.DELETE("/products/:id", product.DeleteTrack)

route.POST("/createuser", user.CreateUserDB)
route.POST("/login", user.Login)
userRoute.GET("/users/:id", user.GetAllUsers)

route.Run()*/
