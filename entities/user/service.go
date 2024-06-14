package user

import (
	"crudgin/pkg/db"
	"crudgin/pkg/utils/jwttoken"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var loginJwt struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginJwt); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user Users
	var err error
	result := db.DB.Where("login = ?", loginJwt.Login).First(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error BD"})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные для входа"})
		return
	}

	// Проверка пароля
	if loginJwt.Password != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные для входа"})
		return
	}

	// Генерация JWT токена
	tokenString, err := jwttoken.GenerateToken(user.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})
}

// список пользователей
func GetAllUsersService(input []Users) error {
	if input == nil {
		err := errors.New("неверный input: поле не может быть пустым")
		log.Println(err)
		return err
	}
	err := GetAllUserDB(input)
	if err != nil {
		log.Println("Ошибка при создании пользователя в базе данных:", err)
		return err
	}
	return err
}

// получение одного пользователя
func GetUserService(input Users) error {
	if input.Login == "" {
		err := errors.New("неверный input: поле Login не может быть пустым")
		log.Println(err)
		return err
	}
	err := GetUserDB(input)
	if err != nil {
		log.Println("Ошибка при создании пользователя в базе данных:", err)
		return err
	}
	return err
}

func CreateUserService(input Users) error {
	if input.Login == "" {
		err := errors.New("неверный input: поле Login не может быть пустым")
		log.Println(err)
		return err
	}
	err := CreateUserDB(input)
	if err != nil {
		log.Println("Ошибка при создании пользователя в базе данных:", err)
		return err
	}
	return err
}

// обновление пользователя
func UpdateUserService(input Users) error {
	err := UpdateUserDB(input)
	if err != nil {
		log.Println("Не удалось обновить пользователя")
		return err
	}
	return nil
}

// удаление пользователя
func DeleteUserService(input Users) error {
	if input.ID == 0 {
		err := errors.New("неверный input: поле ID не может быть пустым")
		log.Println(err)
		return err
	}
	err := DeleteUserDB(input)
	if err != nil {
		log.Println("не удалось удалить пользователя")
		return err
	}
	return err
}
