package router

import "github.com/gin-gonic/gin"

// New generates a gin adapter for handling our routes
func New() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
