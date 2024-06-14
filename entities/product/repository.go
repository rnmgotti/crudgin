package product

import (
	"crudgin/pkg/db"
	"log"
)

func CreateProductDB(input Products) error {
	if err := db.DB.Create(&input).Error; err != nil {
		log.Println("Не удалось создать продукт:", err)
		return err
	}
	return nil
}

func DeleteProductDB(input Products) error {
	if err := db.DB.Where("id=?").First(&input).Error; err != nil {
		log.Println("такого producta нет")
	}
	if err := db.DB.Delete(&input).Error; err != nil {
		log.Println("Не удалось удалить продукт", err)
		return err
	}
	return nil
}

func UpdateProductDB(input Products) error {
	var Product Products
	if err := db.DB.First(&Product, input.ID).Error; err != nil {
		log.Println("Продукт не найден:", err)
		return err
	}

	if err := db.DB.Model(&Product).Updates(input).Error; err != nil {
		log.Println("Не удалось обновить продукт:", err)
		return err
	}
	return nil
}

func GetAllProductDB(input []Products) error {
	if err := db.DB.Find(&input).Error; err != nil {
		log.Println("Не удалось получить список productov")
	}
	return nil
}

func GetProductDB(input Products) error {
	if err := db.DB.Where("login = ?").First(&input).Error; err != nil {
		log.Println("такого producta нет")
	}
	return nil
}
