package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Accounts struct {
	ID      int
	Balance int
	Version int `gorm:"version"`
}

type Transactions struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        int
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

	//db.AutoMigrate(&Accounts{})
	//db.AutoMigrate(&Transactions{})
	//
	//db.Create(&Accounts{Balance: 1000})
	//db.Create(&Accounts{Balance: 200})

	//var a1 Accounts
	//var a2 Accounts
	//db.Where("ID = ?", 1).First(&a1)
	//db.Where("ID = ?", 2).First(&a2)
	//fmt.Println(a1)
	//fmt.Println(a2)

	err := db.Transaction(func(tx *gorm.DB) error {
		var a1 Accounts
		var a2 Accounts
		tx.Where("ID = ?", 1).First(&a1)
		tx.Where("ID = ?", 2).First(&a2)
		fmt.Println(a1)
		fmt.Println(a2)

		a1.Balance -= 100
		if err := tx.Save(&a1).Error; err != nil {
			return err
		}
		if a1.Balance < 0 {
			return errors.New("a1 余额不足")
		}
		a2.Balance += 100

		if err := tx.Save(&a2).Error; err != nil {
			return err
		}
		t := Transactions{
			FromAccountId: a1.ID,
			ToAccountId:   a2.ID,
			Amount:        100,
		}
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
