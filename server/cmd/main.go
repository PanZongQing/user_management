package main

import (
	"server/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config/config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":8080")
}
