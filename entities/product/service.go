package product

import (
	"errors"
	"log"
)

var Err error

func GetAllProductService(input []Products) error {
	if input == nil {
		err := errors.New("неверный input: поле не может быть пустым")
		log.Println(err)
		return err
	}
	err := GetAllProductDB(input)
	if err != nil {
		log.Println("Ошибка при получениии продукта из базе данных:", err)
		return err
	}
	return err
}

func GetProductService(input Products) error {
	if input.Name == "" {
		err := errors.New("неверный input: поле Login не может быть пустым")
		log.Println(err)
		return err
	}
	err := GetProductDB(input)
	if err != nil {
		log.Println("Ошибка при создании пользователя в базе данных:", err)
		return err
	}
	return err
}

func CreateProductService(input Products) error {
	if input.Name == "" {
		err := errors.New("неверный input: поле Login не может быть пустым")
		log.Println(err)
		return err
	}
	err := CreateProductDB(input)
	if err != nil {
		log.Println("Ошибка при создании пользователя в базе данных:", err)
		return err
	}
	return err
}
func UpdateProductService(input Products) error {
	err := UpdateProductDB(input)
	if err != nil {
		log.Println("Не удалось обновить пользователя")
		return err
	}
	return nil
}

func DeleteProductService(input Products) error {
	if input.ID == 0 {
		err := errors.New("неверный input: поле ID не может быть пустым")
		log.Println(err)
		return err
	}
	err := DeleteProductDB(input)
	if err != nil {
		log.Println("не удалось удалить product")
		return err
	}
	return err
}
