package sql

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
}
