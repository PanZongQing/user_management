package router

import (
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	freenas := r.Group("/freenas")
	{
		freenas.GET("/users", controllers.GetFreenasUsers)
		freenas.POST("/users", controllers.CreateFreenasUser)
		freenas.PUT("/users", controllers.UpdateFreenasUser)
		freenas.DELETE("/users", controllers.DeleteFreenasUser)
	}
	{
		freenas.GET("/groups", controllers.GetFreenasGroup)
	}
	gitlab := r.Group("/gitlab")
	{
		gitlab.GET("/users", controllers.GetGitlabUser)
		gitlab.POST("/users", controllers.CreateGitlabUser)
		gitlab.PUT("/users", controllers.UpdateGitlabUser)
		gitlab.DELETE("/users", controllers.DeleteGitlabUser)
	}
	gpu := r.Group("/gpu")
	{
		gpu.GET("/users", controllers.GetGpuUser)
		gpu.POST("/users", controllers.CreateGpuUser)
		gpu.DELETE("/users", controllers.DeleteGpuUser)
	}

}
