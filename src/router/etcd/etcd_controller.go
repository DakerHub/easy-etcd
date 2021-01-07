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
	Key   string   `json:"key" binding:"required"`
}

type PutKvPamras struct {
	Nodes []string `json:"nodes" binding:"required"`
	Key   string   `json:"key" binding:"required"`
	Value string   `json:"value" binding:"required"`
}

func RouteConnect(c *gin.Context) {
	params := GetAllKVParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")

	cli, err := Connect(params.Nodes, auth.(Auth))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	cli.Close()

	c.JSON(200, gin.H{
		"message": "Ok",
	})
}

func RouteGetAllKv(c *gin.Context) {
	params := GetAllKVParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")

	res, err := GetByPrefix(params.Nodes, auth.(Auth), "")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
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

	auth, _ := c.Get("auth")

	err := Del(params.Nodes, auth.(Auth), params.Key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func RoutePutKv(c *gin.Context) {
	params := PutKvPamras{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")
	_, err := Put(params.Nodes, auth.(Auth), params.Key, params.Value)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
