package controllers

import (
	"fmt"
	"net/http"
	"server/internal/models"
	"server/internal/services"

	"github.com/gin-gonic/gin"
)

func GetFreenasUsers(c *gin.Context) {
	username := c.Query("username")
	users, err := services.GetFreenasUsers(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetFreenasGroup(c *gin.Context) {
	groupname := c.Query("group")
	users, err := services.GetFreenasGroup(groupname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateFreenasUser(c *gin.Context) {
	var user models.FreenasUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateFreenasUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func UpdateFreenasUser(c *gin.Context) {
	var user models.FreenasUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateFreenasUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteFreenasUser(c *gin.Context) {
	user := c.Query("username") // 从 URL 参数获取用户名
	fmt.Println(user)
	err := services.DeleteFreenasUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil) // 返回204 No Content
}

// 更新和删除用户的函数类似
