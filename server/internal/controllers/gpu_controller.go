package controllers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func GetGpuUser(c *gin.Context) {
	userinfo := c.Query("username")
	data, err := services.GetGpuUser(userinfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})

}

func CreateGpuUser(c *gin.Context) {
	//TODO
	var user models.GpuUserRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateGpuUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func DeleteGpuUser(c *gin.Context) {
	var user models.GpuUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.DeleteGpuUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除用户成功"})
}
