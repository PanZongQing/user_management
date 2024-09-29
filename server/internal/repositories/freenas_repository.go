package repositories

import (
	"encoding/json"
	"server/internal/models"

	"github.com/go-resty/resty/v2"
)

func GetFreenasUsers() ([]models.FreenasUser, error) {
	client := resty.New()
	resp, err := client.R().Get("http://your_freenas_ip/api/v1/users")
	if err != nil {
		return nil, err
	}
	var users []models.FreenasUser
	if err := json.Unmarshal(resp.Body(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// 其他操作类似
