package models

type FreenasUserRequest struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	GroupID   int    `json:"group"`
	GroupName string `json:"group_name"`
}
type FreenasUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
type FreenasUserUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FreenasGroup struct {
	ID    int    `json:"id"`
	Group string `json:"group"`
}

type GitlabUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type GitlabUserResponse struct{}
