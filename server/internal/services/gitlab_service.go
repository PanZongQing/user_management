package services

import (
	"server/internal/models"
	"server/internal/repositories"

	"github.com/xanzy/go-gitlab"
)

func GetGitlabUser(username string) (*gitlab.User, error) {
	data, err := repositories.GetGitlabUser(username)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func CreateGitlabUser(data models.GitlabUserRequest) error {
	//TODO: implement
	err := repositories.CreateGitlabUser(data)
	if err != nil {
		return err
	}
	return nil
}
func UpdateGitlabUser(data models.GitlabUserRequest) error {
	//TODO: implement
	err := repositories.UpdateGitlabUser(data)
	if err != nil {
		return err
	}
	return nil
}
func DeleteGitlabUser(data models.GitlabUserRequest) error {
	//TODO: implement
	err := repositories.DeleteGitlabUser(data)
	if err != nil {
		return err
	}
	return nil
}
