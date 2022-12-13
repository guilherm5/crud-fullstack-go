package handlers

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}

type Users struct {
	ID    int    `json: "id" gorm:"primaryKey"`
	Nome  string `json: "nome" form: "nome" binding:"required"`
	Email string `json: "email" form: "email binding:"required""`
}
