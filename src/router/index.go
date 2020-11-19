package router

import (
	"easy-etcd/src/router/etcd"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() {
	r := gin.Default()

	r.POST("/api/kv/query", etcd.RouteGetAllKv)
	r.POST("/api/kv/delete", etcd.RouteDeleteKey)
	r.POST("/api/kv/put", etcd.RoutePutKv)

	r.Run("0.0.0.0:9600")
}
