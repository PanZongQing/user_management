package services

import (
	"fmt"
	"server/internal/models"
	"server/internal/repositories"
)

type FreenasUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetFreenasUsers(username string) (*models.FreenasUser, error) {
	users, err := repositories.GetFreenasUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 实现获取 FreeNAS 用户的逻辑

func CreateFreenasUser(user models.FreenasUser) error {
	err := repositories.CreateFreenasUser(user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFreenasUser(username string, updatedUser models.FreenasUser) error {
	existingUser, err := repositories.GetFreenasUserByUsername(username)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return fmt.Errorf("用户 %s 不存在", username)
	}

	err = repositories.UpdateFreenasUser(username, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFreenasUser(username string) error {
	existingUser, err := repositories.GetFreenasUserByUsername(username)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return fmt.Errorf("用户 %s 不存在", username)
	}

	err = repositories.DeleteFreenasUser(username)
	if err != nil {
		return err
	}
	return nil
}
