package product

import (
	"crudgin/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(context *gin.Context) {
	var products []Products
	db.DB.Find(&products)

	context.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProducts(context *gin.Context) {

	var product Products
	if err := db.DB.Where("id=?", context.Param("id")).First(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"products": product})
}

func CreateProducts(context *gin.Context) {
	var input CreateProductsInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := Products{Name: input.Name, Price: input.Price}
	db.DB.Create(&product)

	context.JSON(http.StatusOK, gin.H{"products": product})
}

func UpdateProduct(context *gin.Context) {
	var product Products
	if err := db.DB.Where("id=?", context.Param("id")).First(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}
	var input UpdateProductsInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Update(input)

	context.JSON(http.StatusOK, gin.H{"products": product})
}

func DeleteTrack(context *gin.Context) {
	var product Products
	if err := db.DB.Where("id = ?", context.Param("id")).First(&product).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	db.DB.Delete(&product)

	context.JSON(http.StatusOK, gin.H{"products": true})
}
