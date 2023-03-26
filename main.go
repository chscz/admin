package main

import (
	"admin/internal/route"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

type account struct {
	UserID   string `gorm:"column:user_id"`
	Password string `gorm:"column:PASSWORD"`
	Role     int    `gorm:"column:role"`
}

func main() {
	// db연결
	var acc account
	fmt.Println("[0]")
	DBConn = dbConn()
	fmt.Println("[4]")
	if err := DBConn.Table("account").Find(&acc).Error; err != nil {
		fmt.Println("[5-1]")
		fmt.Println(err)
	} else {
		fmt.Println("[5-2]")
		fmt.Println("UserID: ", acc.UserID)
		fmt.Println("password: ", acc.Password)
		fmt.Println("role: ", acc.Role)
	}
	// Gin 라우터 생성
	fmt.Println("[6]")
	router := route.SetRouter(DBConn)
	fmt.Println("[7]")
	// 서버 시작
	router.Run(":8080")
	fmt.Println("[8]")
}

func dbConn() *gorm.DB {
	fmt.Println("[1]")
	dsn := "root:1111@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("[2]")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("[3]")
	if err != nil {
		fmt.Println("[3-1]")
		panic("db error")
	}
	fmt.Println("[0]")
	return db
}
