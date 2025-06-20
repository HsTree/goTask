package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employees struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

type Employee struct {
	Name   string
	Salary int
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

	//db.AutoMigrate(&Employees{})
	//
	//db.Create([]Employees{
	//	{Name: "t1", Department: "测试部", Salary: 10000},
	//	{Name: "t2", Department: "技术部", Salary: 20000},
	//	{Name: "t3", Department: "技术部", Salary: 30000},
	//})
	var e []Employee
	//db.Debug().Model(&Employees{}).Where("Department = ?", "技术部").Scan(&e)
	//fmt.Println(e)

	db.Debug().Model(&Employees{}).Order("Salary desc").First(&e).Scan(&e)
	fmt.Println(e)
}
