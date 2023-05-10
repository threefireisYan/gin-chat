package main

import (
	"fmt"
	models "ginchat/modles"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})
	// Create
	user := &models.UserBasic{}
	user.Name = "三火"
	db.Create(user)
	// Read
	fmt.Println(db.First(&user, 1))
	db.Model(&user).Update("PassWord", "1234")
}
