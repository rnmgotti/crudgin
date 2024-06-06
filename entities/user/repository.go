package user

import (
	"crudgin/entities/product"
	"crudgin/pkg/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserDB(c *gin.Context) {
	var input Users
	var err error
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := Users{Login: input.Login, Password: input.Password}
	db.DB.Create(&user)
	if err != nil {
		log.Println("Не удалось создать пользователя")
	}

	log.Println(user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func MigrateModels() {
	db.DB.AutoMigrate(&product.Products{})
	db.DB.AutoMigrate(&Users{})
}
