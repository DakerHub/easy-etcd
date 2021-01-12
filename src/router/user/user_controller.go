package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RouteLogin(c *gin.Context) {
	params := LoginParams{}
	if c.ShouldBind(&params) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请求参数错误",
		})
		return
	}

	token, err := Login(params.Username, params.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, false)

	c.JSON(200, gin.H{
		"message": "Ok",
	})
}
