package product

import (
	"crudgin/pkg/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProducts(c *gin.Context) {
	var input Products
	var err error
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		log.Println("Не удалось создать продукт")
	}

	product := Products{Name: input.Name, Price: input.Price}
	db.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"products": product})
}
