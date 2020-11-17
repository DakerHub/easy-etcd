package etcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllKVParams struct {
	Nodes []string `json:"nodes" binding:"required"`
}

type DelKeyPamras struct {
	Nodes []string `json:"nodes" binding:"required"`
	Key   string   `json:"key", binding:"required"`
}

func RouteGetAllKv(c *gin.Context) {
	params := GetAllKVParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	res, err := GetByPrefix(params.Nodes, "")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, UtilProcessGetRes(res))
}

func RouteDeleteKey(c *gin.Context) {
	params := DelKeyPamras{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	err := Del(params.Nodes, params.Key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
