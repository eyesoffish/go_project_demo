package service

import "github.com/gin-gonic/gin"

// Index
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "welcome"})
}
