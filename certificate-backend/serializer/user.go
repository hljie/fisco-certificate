package serializer

import "certificate-backend/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"usename"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
