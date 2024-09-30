package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func GetFreenasUser(username string) (*models.FreenasUser, error) {
	client := resty.New()
	url := viper.GetString("freenas.host" + username)
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	var user models.FreenasUser
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// 其他操作类似
func CreateFreenasUser(user models.FreenasUser) error {
	client := resty.New()
	url := viper.GetString("freenas.host") + "/users" // 确保 URL 正确拼接
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Post(url)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
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
