package handlers

import (
	"crudgin/entities/product"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input product.Products
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := product.CreateProductService(input)
	if err != nil {
		log.Println("Не удалось создать c.AbortWithStatusJSON")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать product"})

	}
	log.Println("Product успешно создан:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Product успешно создан", "product": input})
}

func DeleteProduct(c *gin.Context) {
	var input product.Products
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := product.DeleteProductService(input)
	if err != nil {
		log.Println("Не удалось удалить продукт")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить Продукт"})
	}
	log.Println("Пррдукт успешно удален:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Продукт успешно удален", "product": input})
}

func UpdateProduct(c *gin.Context) {
	var input product.Products
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := product.UpdateProductService(input)
	if err != nil {
		log.Println("Не удалось обновить продукт:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить продукт"})
		return
	}

	log.Println("Продукт успешно обновлен:", input)
	c.JSON(http.StatusOK, gin.H{"message": "Продукт успешно обновлен", "product": input})
}

func GetProduct(c *gin.Context) {
	var input product.Products
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := product.GetProductService(input)
	if err != nil {
		log.Println("Не удалось получить продукт")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить продукт"})

	}
	log.Println("Продукт успешно получен:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Продукт успешно получен", "product": input})
}

func GetAllProduct(c *gin.Context) {
	var input []product.Products
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := product.GetAllProductService(input)
	if err != nil {
		log.Println("Не удалось получить продукт")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить продукты"})

	}
	log.Println("Продукт успешно получен:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Продукт успешно получен", "product": input})
}
