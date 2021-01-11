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

type CreateSnapshotParams struct {
	Nodes []string `json:"nodes" binding:"required"`
}

func RouteCreateSnapshot(c *gin.Context) {
	params := CreateSnapshotParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")

	err := Snapshot(params.Nodes, auth.(Auth))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "创建Snapshot失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}

type GetSnapshotParams struct {
	Nodes []string `json:"nodes" binding:"required"`
}

func RouteGetSnapshot(c *gin.Context) {
	params := GetSnapshotParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")

	_, err := Connect(params.Nodes, auth.(Auth))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	snapshots, err := GetAllLocalSnapshot(params.Nodes[0])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"data":    snapshots,
	})
}

type RestoreParams struct {
	Nodes      []string `json:"nodes" binding:"required"`
	SnapshotID string   `json:"snapshotID" binding:"required"`
	Clear      bool     `json:"clear" binding:"required"`
}

func RouteRestore(c *gin.Context) {
	params := RestoreParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")

	err := RestoreBySnapshot(params.Nodes, auth.(Auth), params.SnapshotID, params.Clear)
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

type DelSnapshotParams struct {
	Nodes      []string `json:"nodes" binding:"required"`
	SnapshotID string   `json:"snapshotID" binding:"required"`
}

func RouteDeleteSnapshot(c *gin.Context) {
	params := DelSnapshotParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	auth, _ := c.Get("auth")

	_, err := Connect(params.Nodes, auth.(Auth))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = DeleteLocalSnapshot(params.SnapshotID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "删除Snapshot失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
