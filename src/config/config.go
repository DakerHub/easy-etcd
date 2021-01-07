package config

import (
	"os"
)

// StaticDir 静态文件存储路径
const StaticDir = "/var/www/html"

// Port 端口号
const Port = "9600"

// Host 主机名
const Host = "0.0.0.0"

// AdminUsername 系统登录用户
var AdminUsername = os.Getenv("EASY_ETCD_USERNAME")

// AdminPassword 系统登录密码
var AdminPassword = os.Getenv("EASY_ETCD_PASSWORD")

// AdminSecret 系统秘钥
var AdminSecret = os.Getenv("Secret")
