package repositories

import (
	"crypto/sha256"
	"encoding/hex"
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

type AuthResquest struct {
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDataResponse struct {
	Data struct {
		AccessToken string `json:"access"`
	} `json:"data"`
}

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
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
func GetToken() (string, error) {
	url := viper.GetString("gpu.url") + "/user/token"
	client := resty.New()
	hashedPassword := HashPassword(gpuPassword)
	authRequest := AuthResquest{
		Admin:    true,
		Username: gpuUser,
		Password: hashedPassword,
	}

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(authRequest).Post(url)
	if err != nil {
		return "", err
	}
	var tokenData TokenDataResponse
	if err := json.Unmarshal([]byte(resp.String()), &tokenData); err != nil {
		return "", err
	}
	accessToken := tokenData.Data.AccessToken
	return accessToken, nil

}

func GetGpuUser(username string) (*models.GpuUser, error) {
	url := viper.GetString("gpu.url") + "/back_admin/business/user"
	client := resty.New()
	token, err := GetToken()
	if err != nil {
		return nil, err
	}
	gettok := "Bearer " + token
	fmt.Printf("token是:%s", gettok)
	resp, err := client.R().SetHeader("Authorization", gettok).SetQueryParam("number", username).Get(url)
	if err != nil {
		return nil, err
	}
	fmt.Printf("响应数据是：%s", resp.String())

	var users models.GpuUserResponse
	if err := json.Unmarshal([]byte(resp.String()), &users); err != nil {
		return nil, err
	}
	fmt.Println(users.Data)

	// 遍历用户列表，查找匹配的用户
	for _, u := range users.Data {
		// fmt.Println(u.Username)
		if u.Username == username { // 假设 FreenasUser 结构体有 Username 字段
			fmt.Printf("找到用户:%s", u.Username)
			return &u, nil // 返回匹配的用户指针
		}
	}

	return nil, fmt.Errorf("用户 %s 未找到", username)
}

func CreateGpuUser(data models.GpuUserRequest) error {
	url := viper.GetString("gpu.url") + "/back_admin/business/user"
	client := resty.New()
	token, err := GetToken()
	if err != nil {
		return err
	}
	_, err = client.R().SetHeader("Authorization", "Bearer "+token).SetBody(data).Post(url)
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
