package sql

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	//这里是注册路由，所有使用到的get、post请求都在这里定义
	r.GET("/save", SaveUser)
	r.GET("/get", GetUser)
	r.GET("/getall", GetAllUser)
	r.GET("/update", UpdateUser)
	r.GET("/delete", DeleteUser)
}
