package router

import (
	"easy-etcd/src/config"
	"easy-etcd/src/router/etcd"
	"easy-etcd/src/router/user"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/login", user.RouteLogin)

	api.Use(etcd.GetAuth())
	{
		api.POST("/connect", etcd.RouteConnect)
		api.POST("/kv/query", etcd.RouteGetAllKv)
		api.POST("/kv/delete", etcd.RouteDeleteKey)
		api.POST("/kv/put", etcd.RoutePutKv)
		api.POST("/snapshot/create", etcd.RouteCreateSnapshot)
		api.POST("/snapshot/query", etcd.RouteGetSnapshot)
		api.POST("/snapshot/delete", etcd.RouteDeleteSnapshot)
		api.POST("/restore", etcd.RouteRestore)
	}

	r.Static("/", config.StaticDir)

	r.Run(config.Host + ":" + config.Port)
}
