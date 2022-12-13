package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

type Users struct {
	ID    int    `json: "id" gorm:"primaryKey"`
	Nome  string `json: "nome" form: "nome" binding:"required"`
	Email string `json: "email" form: "email binding:"required""`
}

func OpenDB(DB *gorm.DB) Handler {

	return Handler{DB}
}

func Init() *gorm.DB {
	dsn := "host=localhost user=# password=# dbname=# port=# sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao realizar conectar com o postgress")
	}
	fmt.Println("Sucesso ao realizar conexão com o postgres")
	db.AutoMigrate(&Users{})
	return db

}

func (h Handler) DeleteUser(c *gin.Context) {
	var user Users
	if err := h.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Invalido ou já foi apagado. Aqui está uma lista de funcionarios atuais e seus ID'S"})
	} else {
		id := c.Params.ByName("id")
		d := h.DB.Where("id = ?", id).Delete(&user)
		fmt.Println(d)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}

func (h Handler) ListUser(c *gin.Context) {
	var listUsers []Users
	h.DB.Find(&listUsers)
	//listUsers = make([]Users, 0)
	listUsers = append(listUsers, Users{})
	c.HTML(http.StatusOK, "listar.html", gin.H{
		"listUsers": listUsers,
	})

}

func (h Handler) Postform(c *gin.Context) {
	var users = []Users{{Nome: c.Request.PostFormValue("nome"), Email: c.Request.PostFormValue("email")}}

	h.DB.Create(&users)

}

func main() {
	db := Init()
	h := OpenDB(db)

	r := gin.Default()
	r.Static("/static", "./public")
	r.Static("/css", "public/css")
	r.LoadHTMLGlob("public/*.html")

	r.GET("/list", h.ListUser)
	r.POST("/upload", h.Postform)
	r.DELETE("/delet/:id", h.DeleteUser)

	r.Run(":7878")
}
