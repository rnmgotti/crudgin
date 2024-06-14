package user

import (
	"crudgin/entities/product"
	"crudgin/pkg/db"
	"log"
)

func MigrateModels() {
	db.DB.AutoMigrate(&product.Products{})
	db.DB.AutoMigrate(&Users{})
}

func CreateUserDB(input Users) error {
	if err := db.DB.Create(&input).Error; err != nil {
		log.Println("Не удалось создать пользователя:", err)
		return err
	}
	return nil
}

func DeleteUserDB(input Users) error {
	if err := db.DB.Where("id=?").First(&input).Error; err != nil {
		log.Println("такого usera нет")
	}
	if err := db.DB.Delete(&input).Error; err != nil {
		log.Println("Не удалось удалить пользователя", err)
		return err
	}
	return nil
}

func UpdateUserDB(input Users) error {
	var user Users
	if err := db.DB.First(&user, input.ID).Error; err != nil {
		log.Println("Пользователь не найден:", err)
		return err
	}

	if err := db.DB.Model(&user).Updates(input).Error; err != nil {
		log.Println("Не удалось обновить пользователя:", err)
		return err
	}
	return nil
}

func GetAllUserDB(input []Users) error {
	if err := db.DB.Find(&input).Error; err != nil {
		log.Println("Не удалось получить список пользователей")
	}
	return nil
}

func GetUserDB(input Users) error {
	if err := db.DB.Where("login = ?").First(&input).Error; err != nil {
		log.Println("такого usera нет")
	}
	return nil
}
