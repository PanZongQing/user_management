package repositories

import (
	"fmt"
	"log"
	"server/internal/models"

	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

var (
	gl    *gitlab.Client
	token string
	url   string
)

// 初始化GitLab客户端
func init() {
	Loadconfig()
	var err error
	token = viper.GetString("gitlab.token")
	url = viper.GetString("gitlab.url")
	gl, err = gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		log.Fatalf("初始化GitLab客户端失败: %v", err)
	}
	fmt.Println("初始化完成")
}

// 获取GitLab用户信息
func GetGitlabUser(username string) (*gitlab.User, error) {
	if username == "" {
		return nil, fmt.Errorf("用户名不能为空")
	}

	// 根据用户名搜索用户
	opts := &gitlab.ListUsersOptions{
		Username: gitlab.Ptr(username),
	}
	users, _, err := gl.Users.ListUsers(opts)
	if err != nil {
		return nil, fmt.Errorf("搜索用户 %s 失败: %v", username, err)
	}

	// 输出找到的用户信息
	if len(users) > 0 {
		user := users[0]
		fmt.Printf("找到用户: %+v\n", user)
	} else {
		fmt.Printf("未找到用户名为 '%s' 的用户。\n", username)
	}
	return users[0], nil
}

// 创建gitlab用户
func CreateGitlabUser(data models.GitlabUserRequest) error {
	user, _, err := gl.Users.CreateUser(&gitlab.CreateUserOptions{
		Username:         gitlab.Ptr(data.Username),
		Password:         gitlab.Ptr(data.Password),
		CanCreateGroup:   gitlab.Ptr(true),
		Name:             gitlab.Ptr(data.FullName),
		Email:            gitlab.Ptr(data.Email),
		SkipConfirmation: gitlab.Ptr(true),
	})
	if err != nil {
		return fmt.Errorf("创建用户%s失败。", user.Username)
	}
	return nil
}

// 修改用户方法
func UpdateGitlabUser(data models.GitlabUserRequest) error {
	user, err := GetGitlabUser(data.Username)
	if err != nil {
		return err
	}
	_, _, err = gl.Users.ModifyUser(user.ID, &gitlab.ModifyUserOptions{
		Username: gitlab.Ptr(data.Username),
		Password: gitlab.Ptr(data.Password),
	})
	if err != nil {
		return fmt.Errorf("更新用户%s失败。", user.Username)
	}
	return nil
}

// 删除用户方法
func DeleteGitlabUser(data models.GitlabUserRequest) error {
	user, err := GetGitlabUser(data.Username)
	if err != nil {
		return err
	}
	_, err = gl.Users.DeleteUser(user.ID)
	if err != nil {
		return fmt.Errorf("删除用户%s失败。", user.Username)
	}
	return nil
}
