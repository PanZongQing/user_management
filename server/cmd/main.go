package main

import (
	"server/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// 配置CORS

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"}, // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,         // 是否允许发送Cookie
		MaxAge:           24 * 60 * 60, // 预检请求的缓存时间
	}))

	router.SetupRouter(r)
	r.Run(":8080")
}
