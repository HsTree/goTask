package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	PostCount int
}

type Post struct {
	ID           uint
	Title        string
	UserID       uint
	User         User
	CommentCount int
}

type Comment struct {
	Text   string
	PostID uint
	Post   Post
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	c := p.User.PostCount + 1
	//fmt.Println("Post BeforeCreate", c)
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("PostCount", c).Error; err != nil {
		return err
	}
	return nil
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
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//u := User{Name: "ttt1"}
	//db.Create(&u)

	//var u User
	//db.Where("ID = ?", 1).First(&u)
	//fmt.Println(u)
	//
	//p := Post{Title: "ttt2", User: u}
	//db.Create(&p)
	//
	//db.Create(&Comment{Text: "ttt2", Post: p})

	var u User
	db.Model(&User{}).Order("post_count desc").First(&u)
	fmt.Println(u)

	var p Post
	db.Model(&Post{}).Order("comment_count desc").First(&p)
	fmt.Println(p)

}
