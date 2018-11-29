package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.POST("/hello", func(c *gin.Context) {
		var msg Msg
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "error:" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + msg.Name,
		})
	})
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name,
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World",
		})
	})
	r.Run(":" + os.Getenv("PORT"))
}
