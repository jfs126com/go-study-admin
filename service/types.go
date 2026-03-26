package service

type LoginPasswordRequest struct {
	Username string `json:"username"` //用户名
	Password string `json:"password"` //登录密码
}
