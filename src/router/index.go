package router

import (
	"easy-etcd/src/config"
	"easy-etcd/src/router/etcd"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() {
	r := gin.Default()

	r.Static("/", config.StaticDir)

	api := r.Group("/api")
	api.Use(etcd.GetAuth())
	{
		api.POST("/connect", etcd.RouteConnect)
		api.POST("/kv/query", etcd.RouteGetAllKv)
		api.POST("/kv/delete", etcd.RouteDeleteKey)
		api.POST("/kv/put", etcd.RoutePutKv)
	}

	r.Run(config.Host + ":" + config.Port)
}
