package services

import (
	"fmt"
	"server/internal/models"
	"server/internal/repositories"
)

func GetFreenasUsers(username string) (*models.FreenasUserResponse, error) {
	users, err := repositories.GetFreenasUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// 获取FreeNAS用户组
func GetFreenasGroup(department_name string) (*models.FreenasGroup, error) {
	groups, err := repositories.GetFreenasGroup(department_name)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// 实现获取 FreeNAS 用户的逻辑

func CreateFreenasUser(user models.FreenasUserRequest) error {
	var group models.FreenasGroup
	switch user.GroupName {
	case "财务部":
		group.Group = "finance"
	case "人事部":
		group.Group = "hr"
	case "技术部":
		group.Group = "it"
	case "AI算法":
		group.Group = "sales"
	case "商务部":
		group.Group = "business"
	case "运营部":
		group.Group = "ops"
	default:
		return fmt.Errorf("部门 %s 不存在", user.GroupName)
	}
	fmt.Println(group.Group)
	id, err := repositories.GetFreenasGroup(string(group.Group))
	if err != nil {
		return err
	}
	fmt.Println(id.ID)
	user.GroupID = id.ID
	err = repositories.CreateFreenasUser(user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFreenasUser(user models.FreenasUserRequest) error {
	existingUser, err := repositories.GetFreenasUserByUsername(user.Username)
	if err != nil {
		return err
	}
	// fmt.Println(&existingUser)
	if existingUser == nil {
		return fmt.Errorf("用户 %s 不存在", user.Username)
	}
	user.ID = existingUser.ID
	var updateJson models.FreenasUserUpdateRequest
	updateJson.Username = user.Username
	updateJson.Password = user.Password

	err = repositories.UpdateFreenasUser(user.ID, updateJson)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFreenasUser(username string) error {
	fmt.Println(username)
	existingUser, err := repositories.GetFreenasUserByUsername(username)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return fmt.Errorf("用户 %s 不存在", username)
	}

	err = repositories.DeleteFreenasUser(existingUser.ID)
	if err != nil {
		return err
	}
	return nil
}
