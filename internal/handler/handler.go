package handler

import (
	"admin/internal/database"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/mysql"
)

func connectDB() *gorm.DB {
	dsn := "root:1111@tcp(127.0.0.1:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db connect error")
	}
	return db
}
func GetTotalUserList() string {
	db := connectDB()
	var userList []*database.User
	if err := db.Find(&userList); err != nil {
		fmt.Println("GetTotalUserList error")
	}
	d, err := json.Marshal(userList)
	if err != nil {
		fmt.Println("json marshal error")
	}
	return string(d)
}

func GetCEOUserList() string {

}

func GetWorkerUserList() string {

}
