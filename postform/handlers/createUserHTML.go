package handlers

import "github.com/gin-gonic/gin"

func (h Handler) CreateUser(c *gin.Context) {
	var users = []Users{{Nome: c.Request.PostFormValue("nome"), Email: c.Request.PostFormValue("email")}}

	h.DB.Create(&users)
}
