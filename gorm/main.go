package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func createUser() {
	// user := User{Name: "张三", Age: 18, Birthday: time.Now()}
	// result := db.Create(user)
	// fmt.Println(user.ID)
	// fmt.Println(result)
}

func main() {
	// 连接数据库
	dsn := "root:tifenbao2014@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	user := User{Name: "张三", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result)
	fmt.Println(err)
}
