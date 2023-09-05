package serializer

import "eletronicMall/model"

// UserVO 用于前端
type UserVO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Type      int    `json:"type"` // 0: 普通用户, 1: 管理员
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户, 用于前端
func BuildUser(user *model.User) *UserVO {
	return &UserVO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Type:      0,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}


