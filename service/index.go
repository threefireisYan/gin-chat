package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回参数
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome!",
	})
}
