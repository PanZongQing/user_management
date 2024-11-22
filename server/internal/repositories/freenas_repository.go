package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/internal/models"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var (
	freenasUser     string
	freenasPassword string
)

func init() {
	Loadconfig()
	var err error
	freenasUser = viper.GetString("freenas.username")
	freenasPassword = viper.GetString("freenas.password")

	if freenasUser == "" || freenasPassword == "" {
		err = fmt.Errorf("freenas username or password not set")
		log.Fatal(err)
	}
}

func Loadconfig() {
	viper.SetConfigFile("../config/config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, ", err)
	}
}

func GetFreenasUserByUsername(username string) (*models.FreenasUserResponse, error) {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/user"

	resp, err := client.R().SetBasicAuth(freenasUser, freenasPassword).Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Println("响应数据是：%s", resp.String())

	var users []models.FreenasUserResponse
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
func CreateFreenasUser(user models.FreenasUserRequest) error {
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
	delete(data, "id")
	updateJson, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %s", err)
	}

	// fmt.Println("Json数据是%s", string(userJson))
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

func UpdateFreenasUser(userid int, user models.FreenasUserUpdateRequest) error {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/user/id/" + strconv.Itoa(userid) // URL 拼接应包含用户名
	resp, err := client.R().SetBasicAuth(freenasUser, freenasPassword).
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Put(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to update user: %s", resp.String())
	}
	return nil
}

func DeleteFreenasUser(userid int) error {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/user/id/" + strconv.Itoa(userid) // URL 拼接应包含用户名
	resp, err := client.R().SetBasicAuth(freenasUser, freenasPassword).Delete(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to delete user: %s", resp.String())
	}
	return nil
}
