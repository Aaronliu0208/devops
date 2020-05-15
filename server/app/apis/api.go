package apis

import "github.com/gin-gonic/gin"

// Ping Ping Pong
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
