package models

import (
	"fmt"
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
	//	1.如果使用动态指定数据库表名
	//err := DB.Scopes(UserTable(user)).Create(&user).Error
	//	2.调用import "gorm.io/gorm"会自动使用已经在gorm.go里面的init函数，
	//result := GetDB().Create(user)
	//	3.可以选择仅插入某些字段
	//result := GetDB().Select("name", "pass_word").Create(user)
	//	4.可以忽略插入某些字段
	//result := GetDB().Omit("author_name", "author_name", "author_email").Create(user)
	//	5.批量插入数据
	////var users = []UserBasic{{Name: "张三"}, {Name: "李四"}}
	//result := GetDB().Omit("author_name", "author_name", "author_email").Create(users)
	//	6.批量-分批插入数据
	////var users = []UserBasic{{Name: "张三"}, {Name: "李四"}}
	//result := GetDB().Omit("author_name", "author_name", "author_email").CreateInBatches(users,2)
	//	7.对密码加密，需要使用map形式进行插入数据
	//result := GetDB().Model(&UserBasic{}).Omit("author_name", "author_name", "author_email").Create(map[string]interface{}{
	//	"name":      "王五",
	//	"pass_word": clause.Expr{SQL: "md5(?)", Vars: []interface{}{"123456"}},
	//})
	//	8.使用原生的sql语句
	//GetDB().Exec("insert into user_basic (name,pass_word) values (?,?)","王五","123456")

	result := GetDB().Omit("author_name", "author_name", "author_email").Create(user)
	affected := result.RowsAffected
	fmt.Println("受影响行数：", affected)
	err := result.Error
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
