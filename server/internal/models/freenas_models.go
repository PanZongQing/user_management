package models

type FreenasUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}

type FreenasGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
