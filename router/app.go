package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
	"log"
)

func Router() *gin.Engine {
	r := gin.Default()
	//service.GetIndex是自定义的返回值方法
	r.GET("/index", service.GetIndex)
	err := r.Run(":8080")
	if err != nil {
		log.Printf("运行失败！")
	}
	return r
}

func InitRouter() {

}
