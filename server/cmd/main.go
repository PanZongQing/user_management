package main

import (
	"fmt"
	"server/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, ", err)
	}
	url := viper.GetString("freenas.host")
	fmt.Println("url:", url)

	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":8080")
}
