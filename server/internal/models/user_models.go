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

type GitlabUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type GpuUserRequest struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

// 定义用户数据结构
type GpuUser struct {
	ID             int     `json:"id"`
	UserType       int     `json:"usertype"`
	Username       string  `json:"username"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	StorageSize    int     `json:"storage_size"`
	Desc           string  `json:"desc"`
	IsActive       bool    `json:"is_active"`
	GroupName      string  `json:"group_name"`
	DateJoined     string  `json:"date_joined"`
	CreateUsername string  `json:"create_username"`
	Name           *string `json:"name"`    // 使用指针以支持 null
	College        *string `json:"college"` // 使用指针以支持 null
	Subject        *string `json:"subject"` // 使用指针以支持 null
	Organization   string  `json:"organization"`
	ValidTime      *string `json:"valid_time"`   // 使用指针以支持 null
	InvalidTime    *string `json:"invalid_time"` // 使用指针以支持 null
}

// 定义整体响应结构
type GpuUserResponse struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	Data     []GpuUser `json:"data"`
	Count    int       `json:"count"`
	Next     *string   `json:"next"`     // 使用指针以支持 null
	Previous *string   `json:"previous"` // 使用指针以支持 null
}
