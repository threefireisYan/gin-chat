package sql

import (
	models "ginchat/modles"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SaveUser(c *gin.Context) {
	user := &models.UserBasic{
		Model:         gorm.Model{},
		Name:          "Dav",
		PassWord:      "123456",
		Phone:         "13222222222",
		Email:         "13222222222@163.com",
		Identity:      "",
		ClientIp:      "",
		ClientPort:    "",
		LoginTime:     0,
		HeartbeatTime: 0,
		LogOutTime:    0,
		IsLogout:      false,
		DeviceInfo:    "",
	}
	//调用models中 user中的保存函数
	models.SaveUser(user)
	//返回新增的数据
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	//调用models中 user中的查询函数
	user := models.GetUserId(5)
	//返回查询的数据
	c.JSON(http.StatusOK, user)
}

func GetAllUser(c *gin.Context) {
	//调用models中 user中的查询函数
	user := models.GetAllUsers()
	//返回查询的数据
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	models.UpdateUser(2)
	//调用models中 user中的查询函数
	user := models.GetUserId(2)
	//返回查询的数据
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	models.DeleteUser(4)
	//调用models中 user中的查询函数
	user := models.GetUserId(4)
	//返回查询的数据
	c.JSON(http.StatusOK, user)
}
