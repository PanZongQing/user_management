package main

import (
	"fmt"
	"server/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Loadconfig() {
	viper.SetConfigFile("../config/config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, ", err)
	}
}

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
