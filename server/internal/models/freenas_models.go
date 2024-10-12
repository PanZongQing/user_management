package models

type FreenasUser struct {
	// ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	GroupID   int    `json:"group"`
	GroupName string `json:"group_name"`
}

type FreenasGroup struct {
	ID    int    `json:"id"`
	Group string `json:"group"`
}
