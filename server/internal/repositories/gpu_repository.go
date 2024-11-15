package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"server/internal/models"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var (
	gpuUser     string
	gpuPassword string
)

func init() {
	Loadconfig()
	var err error
	gpuUser = viper.GetString("gpu.username")
	gpuPassword = viper.GetString("gpu.password")

	if gpuUser == "" || gpuPassword == "" {
		err = fmt.Errorf("GPU username or password not set")
		log.Fatal(err)
	}
}

func GetGpuUser(username string) (*models.GpuUserResponse, error) {
	url := viper.GetString("gpu.url") + "/user/" + username
	client := resty.New()
	resp, err := client.R().SetBasicAuth(gpuUser, gpuPassword).Get(url)
	if err != nil {
		return nil, err
	}
	// fmt.Println("响应数据是：%s", resp.String())

	var users []models.GpuUserResponse
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

func CreateGpuUser(data models.GpuUserRequest) error {
	url := viper.GetString("gpu.url") + "/user"
	client := resty.New()
	_, err := client.R().SetBasicAuth(gpuUser, gpuPassword).SetBody(data).Post(url)
	if err != nil {
		return err
	}
	// fmt.Println("响应数据是：%s", resp.String())
	return nil
}

func DeleteGpuUser(userid int) error {
	url := viper.GetString("gpu.url") + "/user/" + fmt.Sprintf("%d", userid)
	client := resty.New()
	_, err := client.R().SetBasicAuth(gpuUser, gpuPassword).Patch(url)
	if err != nil {
		return err
	}
	// fmt.Println("响应数据是：%s", resp.String())
	return nil
}
