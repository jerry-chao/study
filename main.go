package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* {
    "target": ["jerry1"],
    "target_type": "chat",
    "appkey": "{{org_name}}#{{app_name}}",
    "route_online": true,
    "NS": "notify",
    "sync_device":true,
    "timestamp": 1672885923969,
    "payloads": []
} */

type JsonValue struct {
	Target []string `json:"target" binding:"required"`
	TargetType string `json:"target_type" binding:"required"`
	AppKey string `json:"appkey" binding:"required"`
	RouteOnline string `json:"route_online" binding:"required"`
	NS string `json:"NS" binding:"required"`
	SyncDevice string `json:"sync_device" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/api/easemob.com/:org/:app/notify", func(c *gin.Context) {
		var requestData JsonValue
    	c.BindJSON(&requestData)
		c.JSON(http.StatusOK, gin.H{
			"result": "suc",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run("0.0.0.0:9091") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
