package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=crudgin password=rnmgotti sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	sqlDB := DB
    if err != nil {
        log.Fatalf("Ошибка получения объекта sql.DB: %v", err)
    }

    // Закрытие соединения в конце работы
    defer func() {
        if err := sqlDB.Close(); err != nil {
            log.Fatalf("Ошибка закрытия соединения с базой данных: %v", err)
        }
    }()
}
