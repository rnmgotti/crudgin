package user

import (
	"crudgin/pkg/db"
	"crudgin/pkg/utils/jwttoken"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var loginJwt struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginJwt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user Users
	result := db.DB.Where("login = ?", loginJwt.Login).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error BD"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные для входа"})
		return
	}

	// Проверка пароля
	if loginJwt.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные для входа"})
		return
	}

	// Генерация JWT токена
	tokenString, err := jwttoken.GenerateToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})
}

// список пользователей
func GetAllUsers(context *gin.Context) {
	var users []Users
	db.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{"users": users})
}

// получение одного пользователя
func GetUsers(context *gin.Context) {
	var user Users
	if err := db.DB.Where("id=?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пользователя не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"products": user})
}

// обновление пользователя
func UpdateUser(context *gin.Context) {
	var user Users
	if err := db.DB.Where("id=?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пользователя не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}

// удаление пользователя
func DeleteUser(context *gin.Context) {
	var user Users
	if err := db.DB.Where("id=?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пользователя не существует"})
		return
	}

	db.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"user": true})
}
