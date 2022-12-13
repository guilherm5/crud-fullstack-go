package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) ListUser(c *gin.Context) {
	var listUsers []Users
	h.DB.Find(&listUsers)
	//listUsers = make([]Users, 0)
	listUsers = append(listUsers, Users{})
	c.HTML(http.StatusOK, "listar.html", gin.H{
		"listUsers": listUsers,
	})

}
