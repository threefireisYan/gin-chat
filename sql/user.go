package sql

import (
	models "ginchat/modles"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	models.SaveUser(user)

}
