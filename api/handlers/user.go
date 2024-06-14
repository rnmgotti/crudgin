package handlers

import (
	"crudgin/entities/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input user.Users
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.CreateUserService(input)
	if err != nil {
		log.Println("Не удалось создать пользователя")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})

	}
	log.Println("Пользователь успешно создан:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Пользователь успешно создан", "user": input})
}

func DeleteUSer(c *gin.Context) {
	var input user.Users
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.DeleteUserService(input)
	if err != nil {
		log.Println("Не удалось удалить пользователя")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пользователя"})
	}
	log.Println("Пользователь успешно удален:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален", "user": input})
}

func UpdateUser(c *gin.Context) {
	var input user.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.UpdateUserDB(input)
	if err != nil {
		log.Println("Не удалось обновить пользователя:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить пользователя"})
		return
	}

	log.Println("Пользователь успешно обновлен:", input)
	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно обновлен", "user": input})
}

func GetUser(c *gin.Context) {
	var input user.Users
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.GetUserService(input)
	if err != nil {
		log.Println("Не удалось получить пользователя")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить пользователя"})

	}
	log.Println("Пользователь успешно получен:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Пользователь успешно получен", "user": input})
}

func GetAllUser(c *gin.Context) {
	var input []user.Users
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.GetAllUsersService(input)
	if err != nil {
		log.Println("Не удалось получить продукт")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить пользователя"})

	}
	log.Println("Пользователь успешно получен:", input)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Пользователь успешно получен", "user": input})
}
