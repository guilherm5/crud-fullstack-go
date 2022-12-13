package main

import (
	"crud-test/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	h := database.OpenDB(db)

	r := gin.Default()
	r.Static("/static", "./public")
	r.Static("/css", "public/css")
	r.LoadHTMLGlob("public/*.html")

	r.GET("/list", h.ListUser)
	r.POST("/upload", h.CreateUser)
	r.DELETE("/delet/:id", h.DeleteUser)

	r.Run(":7878")
}
