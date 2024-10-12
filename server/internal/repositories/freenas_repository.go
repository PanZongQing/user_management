package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/internal/models"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func GetFreenasUserByUsername(username string) (*models.FreenasUser, error) {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/user"
	freenasUser := viper.GetString("freenas.username")
	freenasPassword := viper.GetString("freenas.password")
	resp, err := client.R().SetBasicAuth(freenasUser, freenasPassword).Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Println(resp.String())

	var users []models.FreenasUser
	if err := json.Unmarshal(resp.Body(), &users); err != nil {
		return nil, err
	}

	// 遍历用户列表，查找匹配的用户
	for _, u := range users {
		// fmt.Println(u.Username)
		if u.Username == username { // 假设 FreenasUser 结构体有 Username 字段
			return &u, nil // 返回匹配的用户指针
		}
	}

	return nil, fmt.Errorf("用户 %s 未找到", username)
}

func GetFreenasGroup(department_name string) (*models.FreenasGroup, error) {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/group"
	freenasUser := viper.GetString("freenas.username")
	freenasPassword := viper.GetString("freenas.password")
	resp, err := client.R().SetBasicAuth(freenasUser, freenasPassword).Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Println(resp.String())

	var group []models.FreenasGroup
	if err := json.Unmarshal(resp.Body(), &group); err != nil {
		return nil, err
	}

	// 遍历用户列表，查找匹配的用户
	for _, g := range group {
		// fmt.Println(g.Group)
		// fmt.Println(department_name)
		if g.Group == department_name {
			return &g, nil // 返回匹配的用户指针
		}

	}
	return nil, fmt.Errorf("部门 %s 未找到", department_name)
}

// 其他操作类似
func CreateFreenasUser(user models.FreenasUser) error {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/user" // 确保 URL 正确拼接
	userJson, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %s", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(userJson), &data)
	if err != nil {
		log.Fatal(err)
	}
	delete(data, "group_name")
	updateJson, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %s", err)
	}

	// fmt.Println("Json数据是%s", string(userJson))
	freenasUser := viper.GetString("freenas.username")
	freenasPassword := viper.GetString("freenas.password")
	resp, err := client.R().SetBasicAuth(freenasUser, freenasPassword).
		SetHeader("Content-Type", "application/json").
		SetBody(updateJson).
		Post(url)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode())
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to create user: %s", resp.String())
	}
	return nil
}

func UpdateFreenasUser(username string, user models.FreenasUser) error {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/users/" + username // URL 拼接应包含用户名
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Patch(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to update user: %s", resp.String())
	}
	return nil
}

func DeleteFreenasUser(username string) error {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/users/" + username // URL 拼接应包含用户名
	resp, err := client.R().Delete(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("failed to delete user: %s", resp.String())
	}
	return nil
}
