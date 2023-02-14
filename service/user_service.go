package service

import (
	"go_chat/models"

	"github.com/gin-gonic/gin"
)

// Users
// @Tags 首页
// @Success 200
// @Router /users [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
