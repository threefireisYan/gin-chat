package router

import (
	"ginchat/service"
	"ginchat/sql"
	"github.com/gin-gonic/gin"
	"log"
)

func Router() *gin.Engine {
	r := gin.Default()
	//service.GetIndex是自定义的返回值方法
	r.GET("/index", service.GetIndex)
	InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Printf("运行失败！")
	}
	return r
}

// 初始化路由
func InitRouter(r *gin.Engine) {
	sql.RegisterRouter(r)
}
