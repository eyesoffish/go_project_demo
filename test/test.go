package main

import (
	"fmt"
	"go_chat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 如果没有这个表会自动创阿金
	db.AutoMigrate(&models.UserBasic{})

	// create
	user := &models.UserBasic{}
	user.Name = "fishofeyes"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // 根据整形主键查找

	// update
	db.Model(&user).Update("Phone", "17600176001")

	// update 跟新多个字段
	// db.Model(&user).Updates(models.UserBasic{Email: "fishofeyes@gmail.com", Password: "123456"})
	// db.Model(&user).Updates(map[string]interface{}{"Phone": "13312345670", "Email": "77@qq.com"})

	// // 删除
	// db.Delete(&user, 1)
}
