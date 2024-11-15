package services

import (
	"server/internal/models"
	"server/internal/repositories"
)

func GetGpuUser(username string) (*models.GpuUserResponse, error) {
	data, err := repositories.GetGpuUser(username)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func CreateGpuUser(data models.GpuUserRequest) error {
	//TODO: implement
	err := repositories.CreateGpuUser(data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGpuUser(data models.GpuUserResponse) error {
	//TODO: implement
	err := repositories.DeleteGpuUser(data.ID)
	if err != nil {
		return err
	}
	return nil
}
