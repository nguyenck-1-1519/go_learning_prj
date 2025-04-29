package main

import (
	"database/sql"
	"fmt"
	"time"

	"example.com/go-learning-prj/api/route"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mySqlDbType     = "mysql"
	openSqlDbScript = "root:Aa@123456@tcp(localhost:3306)/book_management"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open or Keep connection to DB FAILED", r)
		}
	}()

	db, err := sql.Open(mySqlDbType, openSqlDbScript)
	if err != nil {
		panic(fmt.Sprintf("open db failed %s", err))
	}

	// Check connection before query
	if err := db.Ping(); err != nil {
		panic("keep connection to db failed")
	}

	timeout := time.Duration(120) * time.Second

	http := gin.Default()

	route.Setup(timeout, db, http)

	err = http.Run()
	if err != nil {
		return
	}
}
