package models

import (
	"gorm.io/gorm"
	"log"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogOutTime    uint64
	IsLogout      bool
	DeviceInfo    string
	Admin         bool   `gorm:"-"`                               //这里是不存在于数据库的字段，使用-进行忽略数据库操作
	Author        Author `gorm:"embedded;embeddedPrefix:author_"` //这里是引入嵌套结构体embedded，并设置了前缀embeddedPrefix:author_
}

type Author struct {
	Name  string
	Email string
}

// 这里指明操作表名（重写了gorm中的TableName方法，如果不做重写那么将自动使用UserBasic的表）
func (table *UserBasic) TableName() string {
	return "user_basic"
}

//
//func UserTable(user UserBasic) func(tx *gorm.DB) *gorm.DB {
//	return func(tx *gorm.DB) *gorm.DB {
//		if user.Admin {
//			return tx.Table("admin_users")
//		}
//
//		return tx.Table("users")
//	}
//}

// 新增数据
func SaveUser(user *UserBasic) {
	//如果使用动态指定数据库表名
	//err := DB.Scopes(UserTable(user)).Create(&user).Error
	//	调用import "gorm.io/gorm"会自动使用已经在gorm.go里面的init函数，
	err := GetDB().Create(user).Error
	if err != nil {
		log.Printf("insert user error :", err)
	}
}

// 查询数据
func GetUserId(id int64) UserBasic {
	var user UserBasic
	//	调用import "gorm.io/gorm"会自动使用已经在gorm.go里面的init函数，
	err := GetDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("query user error :", err)
	}
	return user
}

func GetAllUsers() []UserBasic {
	var user []UserBasic

	err := GetDB().Find(&user).Error
	if err != nil {
		log.Printf("query user error :", err)
	}
	return user
}

func UpdateUser(id int64) {
	err := GetDB().Model(&UserBasic{}).Where("id = ?", id).Update("name", "lisi").Error
	if err != nil {
		log.Printf("update user error :", err)
	}
}

func DeleteUser(id int64) {
	//	注意！直接使用delete删除是软删除，仅在数据库记录deleted_at时间，实际数据库数据未删除.如需要使用物理删除数据需要加入Unscoped()方法忽略软删除逻辑
	err := GetDB().Unscoped().Model(&UserBasic{}).Where("id = ?", id).Delete(&UserBasic{}).Error
	//err := DB.Model(&UserBasic{}).Where("id = ?", id).Delete(&UserBasic{}).Error
	if err != nil {
		log.Printf("update user error :", err)
	}
}
