package main

import (
	"fmt"
	"time"

	"example.com/go-learning-prj/api/route"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	mySqlDbType = "mysql"
	mySqlDsn    = "root:Aa@123456@tcp(localhost:3306)/book_management"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open or Keep connection to DB FAILED", r)
		}
	}()

	db, err := gorm.Open(mysql.Open(mySqlDsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open db failed %s", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	timeout := time.Duration(120) * time.Second

	http := gin.Default()

	route.Setup(timeout, db, http)

	err = http.Run()
	if err != nil {
		return
	}
}
