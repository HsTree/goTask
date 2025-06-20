package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type students struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func connect() *gorm.DB {
	dsn := "root:123@tcp(127.0.0.1:3306)/goTask?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func main() {

	db := connect()
	// 迁移数据库
	//err := db.AutoMigrate(&students{})
	//if err != nil {
	//	return
	//}

	// 插入数据
	//s := students{
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//}
	//db.Create(&s)

	// 查询18岁以上的数据
	sl := []students{}
	db.Debug().Where("Age > ?", 10).Find(&sl)
	fmt.Println(sl)

	//更新为 "四年级"
	//s := students{}
	//db.Model(&students{}).Where("Name = ?", "张三").Updates(map[string]interface{}{"Grade": "四年级"})
	//fmt.Println(s)

	// 删除15岁以下
	//s := &students{
	//	Name:  "李四",
	//	Age:   16,
	//	Grade: "一年级",
	//}
	//db.Create(s)
	//
	//db.Where("Age < ?", 15).Unscoped().Delete(&students{})

}
