package database

import (
	"crud-test/handlers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(DB *gorm.DB) handlers.Handler {
	return handlers.Handler{DB}
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=ivaneteJC dbname=postform port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao realizar conectar com o postgress")
	}
	fmt.Println("Sucesso ao realizar conexão com o postgres")
	db.AutoMigrate(&handlers.Users{})
	return db

}
