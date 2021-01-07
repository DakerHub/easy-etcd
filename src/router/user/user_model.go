package user

import (
	"easy-etcd/src/config"
	"easy-etcd/src/jwt"
	"errors"
)

// Login 登录
func Login(username string, password string) (string, error) {
	if username == config.AdminUsername && password == config.AdminPassword {
		return jwt.EncodeJWT(username, password)
	}

	return "", errors.New("用户名密码错误")
}
