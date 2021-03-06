package etcd

import (
	"easy-etcd/src/jwt"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/clientv3"
)

const (
	HEADER_USERNAME = "X-Etcd-Username"
	HEADER_PASSWORD = "X-Etcd-Password"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Res struct {
	Count int64      `json:"count"`
	Kvs   []KeyValue `json:"kvs"`
}

func UtilProcessGetRes(res *clientv3.GetResponse) *Res {
	newRes := &Res{
		Count: res.Count,
		Kvs:   make([]KeyValue, len(res.Kvs)),
	}

	for idx, kv := range res.Kvs {
		newRes.Kvs[idx] = KeyValue{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
	}

	return newRes
}

func GetAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		u, p, err := DecodeBasicAuthHeader(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效用户信息",
			})
			return
		}

		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "请先登录",
			})
			return
		}
		_, _, tokenErr := jwt.DecodeJWT(token)
		if tokenErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token无效，请重新登录",
			})
			return
		}

		c.Set("auth", Auth{
			Username: u,
			Password: p,
		})
		c.Next()
	}
}

func DecodeBasicAuthHeader(header string) (string, string, error) {
	var code string
	parts := strings.SplitN(header, " ", 2)
	if len(parts) == 2 && parts[0] == "Basic" {
		code = parts[1]
	}

	decoded, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", "", err
	}

	auth := strings.SplitN(string(decoded), ":", 2)
	if len(auth) != 2 {
		return "", "", errors.New("Invalid basic auth header")
	}

	return auth[0], auth[1], nil
}
