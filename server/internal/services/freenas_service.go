package services

type FreenasUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetFreenasUsers() ([]FreenasUser, error) {
	// 实现获取 FreeNAS 用户的逻辑
}

func CreateFreenasUser(user FreenasUser) error {
	// 实现创建 FreeNAS 用户的逻辑
}
