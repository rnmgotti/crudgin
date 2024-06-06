package product

import (
	"crudgin/pkg/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Err error

func GetAllProducts(c *gin.Context) {
	var products []Products
	var err error
	db.DB.Find(&products)
	if err != nil {
		log.Println("Не удалось получить список продуктов")
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProducts(c *gin.Context) {

	var product Products
	var err error
	if err := db.DB.Where("id=?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	if err != nil {
		log.Println("Не удалось получить продукт")
	}
	c.JSON(http.StatusOK, gin.H{"products": product})
}

func UpdateProduct(c *gin.Context) {
	var product Products
	var err error
	if err := db.DB.Where("id=?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	if err != nil {
		log.Println("такого продукта нет")
	}
	var input Products
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		log.Println("Не удалось обновить продукт")
	}

	db.DB.Model(&product).Update(input)

	c.JSON(http.StatusOK, gin.H{"products": product})
}

func DeleteTrack(c *gin.Context) {
	var product Products
	var err error
	if err := db.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	if err != nil {
		log.Println("Не удалось удалить продукт")
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"products": true})
}
