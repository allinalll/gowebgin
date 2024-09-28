package database

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func Init() {
	var err error
	Db, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/gomarket?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
		return
	}
}
