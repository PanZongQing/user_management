package router

import (
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	freenas := r.Group("/freenas")
	{
		freenas.GET("/users", controllers.GetUsers)
		freenas.POST("/users", controllers.CreateUser)
		freenas.PUT("/users/:id", controllers.UpdateUser)
		freenas.DELETE("/users/:id", controllers.DeleteUser)
	}
	{
		freenas.GET("/groups", controllers.GetGroup)
	}
	gitlab := r.Group("/gitlab")
	{
		gitlab.GET("/users", controllers.GetUsers)
		gitlab.POST("/users", controllers.CreateUser)
		gitlab.PUT("/users/:id", controllers.UpdateUser)
		gitlab.DELETE("/users/:id", controllers.DeleteUser)
	}

}
