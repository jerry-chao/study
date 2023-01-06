package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/easemob.com/:org/:app/notify", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "suc",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run("0.0.0.0:9091") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
